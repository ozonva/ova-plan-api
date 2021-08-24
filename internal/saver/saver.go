package saver

import (
	"errors"
	"fmt"
	"github.com/ozonva/ova-plan-api/internal/flusher"
	"github.com/ozonva/ova-plan-api/internal/models"
	"sync"
	"time"
)

var (
	BufferIsFull = errors.New("entity buffer is full")
	Terminated   = errors.New("saver is terminated")
)

type Saver interface {
	Save(entity models.Plan) error
	Close() error
}

// NewSaver returns Saver with periodical saving support.
// Function starts goroutine which flushing entities added by Save function with flushInterval period
func NewSaver(
	capacity uint,
	flusher flusher.Flusher,
	flushInterval time.Duration,
) Saver {
	saver := &saver{
		flusher:           flusher,
		buffer:            make([]models.Plan, 0, capacity),
		loopTerminateChan: make(chan struct{}),
	}

	go flushLoop(saver, flushInterval)
	return saver
}

type saver struct {
	sync.Mutex
	onceInit          sync.Once
	buffer            []models.Plan
	loopTerminateChan chan struct{}
	flusher           flusher.Flusher
	terminated        bool
}

func (s *saver) Save(entity models.Plan) error {
	if s.terminated {
		return fmt.Errorf("unable to save: %w", Terminated)
	}

	s.Lock()
	defer s.Unlock()

	if len(s.buffer) == cap(s.buffer) {
		return fmt.Errorf("%w. Unable to save plan with id %v", BufferIsFull, entity.Id)
	}

	s.buffer = append(s.buffer, entity)
	return nil
}

func flushLoop(s *saver, flushInterval time.Duration) {
	flushTicker := time.NewTicker(flushInterval)
	for {
		select {
		case <-flushTicker.C:
			s.flush()
		case <-s.loopTerminateChan:
			return
		}
	}
}

// Close stops flushing
func (s *saver) Close() error {
	if s.terminated {
		return fmt.Errorf("unable to close: %w", Terminated)
	}

	s.loopTerminateChan <- struct{}{}

	s.terminated = true
	s.flush()

	close(s.loopTerminateChan)
	return nil
}

func (s *saver) flush() {
	s.Lock()
	defer s.Unlock()

	failed := s.flusher.Flush(s.buffer)
	s.buffer = make([]models.Plan, 0, cap(s.buffer))
	s.buffer = append(s.buffer, failed...)
}
