package saver

import (
	"context"
	"errors"
	"fmt"
	"github.com/ozonva/ova-plan-api/internal/flusher"
	"github.com/ozonva/ova-plan-api/internal/models"
	"sync"
	"time"
)

var (
	BufferIsFull        = errors.New("entity buffer is full")
	Terminated          = errors.New("saver is terminated")
	NonPositiveArgument = errors.New("argument must be greater than zero")
)

type Saver interface {
	Save(entity models.Plan) error
	Close() error
}

// NewSaver returns Saver with periodical saving support.
//  it starts goroutine which flushing entities added by Save function with flushInterval period.
//  For stop call Close method
func NewSaver(
	ctx context.Context,
	capacity uint,
	flusher flusher.Flusher,
	flushInterval time.Duration,
) (Saver, error) {
	saver := &saver{
		flusher:           flusher,
		buffer:            make([]models.Plan, 0, capacity),
		loopTerminateChan: make(chan struct{}),
		Mutex:             &sync.Mutex{},
		ctx:               ctx,
	}

	if flushInterval < 1 {
		return nil, fmt.Errorf("%w: flushInterval", NonPositiveArgument)
	}

	if capacity < 1 {
		return nil, fmt.Errorf("%w: capacity", NonPositiveArgument)
	}

	//wg guarantees that goroutine has started before function return
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go flushLoop(saver, flushInterval, wg)
	wg.Wait()

	return saver, nil
}

type saver struct {
	*sync.Mutex
	buffer            []models.Plan
	loopTerminateChan chan struct{}
	flusher           flusher.Flusher
	terminated        bool
	ctx               context.Context
}

func (s *saver) Save(entity models.Plan) error {
	if s.terminated {
		return fmt.Errorf("unable to save: %w", Terminated)
	}

	s.Lock()
	defer s.Unlock()

	if len(s.buffer) == cap(s.buffer) {
		s.flush()
	}
	// if there were errors while saving whole buffered elements
	if len(s.buffer) == cap(s.buffer) {
		return fmt.Errorf("%w. Unable to save plan with id %v", BufferIsFull, entity.Id)
	}

	s.buffer = append(s.buffer, entity)
	return nil
}

func flushLoop(s *saver, flushInterval time.Duration, wg *sync.WaitGroup) {
	flushTicker := time.NewTicker(flushInterval)
	wg.Done()
	for {
		select {
		case <-flushTicker.C:
			s.flushWithLock()
		case _, ok := <-s.loopTerminateChan:
			if !ok {
				flushTicker.Stop()
				return
			}
		}
	}
}

// Close stops flushing
func (s *saver) Close() error {
	if s.terminated {
		return fmt.Errorf("unable to close: %w", Terminated)
	}

	close(s.loopTerminateChan)
	s.terminated = true
	s.flushWithLock()

	return nil
}

func (s *saver) flushWithLock() {
	s.Lock()
	defer s.Unlock()

	s.flush()
}

func (s *saver) flush() {
	if len(s.buffer) == 0 {
		return
	}
	failed := s.flusher.Flush(s.ctx, s.buffer)
	s.buffer = make([]models.Plan, 0, cap(s.buffer))
	s.buffer = append(s.buffer, failed...)
}
