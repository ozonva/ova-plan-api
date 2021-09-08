package saver_test

import (
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-plan-api/internal/mocks"
	"github.com/ozonva/ova-plan-api/internal/models"
	"github.com/ozonva/ova-plan-api/internal/saver"
	"math/rand"
	"sync"
	"time"
)

var _ = Describe("Saver", func() {

	var (
		mockController *gomock.Controller
		mockFlusher    *mocks.MockFlusher
		capacity       uint
		flushInterval  time.Duration
		testData       = []models.Plan{
			testPlan(1),
			testPlan(2),
			testPlan(3),
			testPlan(4),
			testPlan(5),
		}
		ctx context.Context
	)

	BeforeEach(func() {
		ctx = context.TODO()
		mockController = gomock.NewController(GinkgoT())
		mockFlusher = mocks.NewMockFlusher(mockController)
	})

	AfterEach(func() {
		mockController.Finish()
	})

	Context("With ordinary arguments", func() {
		JustBeforeEach(func() {
			capacity = 2
			flushInterval = time.Millisecond * 300
		})

		It("Should save one entity in flushing loop", func() {
			var wg = &sync.WaitGroup{}
			wg.Add(1)

			mockFlusher.EXPECT().Flush(gomock.Any(), gomock.Any()).Times(1).
				DoAndReturn(func(ctx2 context.Context, plans []models.Plan) []models.Plan {
					wg.Done()
					return nil
				})

			svr, err := saver.NewSaver(ctx, capacity, mockFlusher, flushInterval)
			Expect(err).Should(BeNil())

			Expect(svr.Save(testData[0])).Should(BeNil())
			//waiting for flush in loop
			wg.Wait()

			Expect(svr.Close()).Should(BeNil())
		})

		It("Should save entities more than batch size", func() {
			defer GinkgoRecover()

			var wg = &sync.WaitGroup{}
			wg.Add(1)
			mockFlusher.EXPECT().Flush(gomock.Any(), testData[:2]).Times(1).
				DoAndReturn(func(ctx2 context.Context, plans []models.Plan) []models.Plan {
					wg.Done()
					return nil
				})
			mockFlusher.EXPECT().Flush(gomock.Any(), testData[2:3]).Times(1).
				DoAndReturn(func(ctx2 context.Context, plans []models.Plan) []models.Plan {
					wg.Done()
					return nil
				})

			svr, err := saver.NewSaver(ctx, capacity, mockFlusher, flushInterval)
			Expect(err).Should(BeNil())

			Expect(svr.Save(testData[0])).Should(BeNil())
			Expect(svr.Save(testData[1])).Should(BeNil())
			//waiting for first flush in loop
			wg.Wait()
			wg.Add(1)
			Expect(svr.Save(testData[2])).Should(BeNil())
			//waiting for second flush in loop
			wg.Wait()

			Expect(svr.Close()).Should(BeNil())
		})
	})

	Context("With giant flush interval", func() {
		JustBeforeEach(func() {
			capacity = 2
			flushInterval = time.Hour * 8800
		})

		It("Should flush entities on close", func() {
			defer GinkgoRecover()

			mockFlusher.EXPECT().Flush(ctx, testData[:1]).Return(nil).Times(1)

			svr, err := saver.NewSaver(ctx, capacity, mockFlusher, flushInterval)
			Expect(err).Should(BeNil())

			Expect(svr.Save(testData[0])).Should(BeNil())
			Expect(svr.Close()).Should(BeNil())
		})

	})
})

func testPlan(id uint64) models.Plan {
	pl := models.NewPlan(id, rand.Uint64(), fmt.Sprintf("test %v", id), "test", time.Now(), time.Now().Add(time.Hour))
	return *pl
}
