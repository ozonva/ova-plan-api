package flusher_test

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozonva/ova-plan-api/internal/flusher"
	"github.com/ozonva/ova-plan-api/internal/mocks"
	"github.com/ozonva/ova-plan-api/internal/models"
	"math/rand"
	"strconv"
	"time"
)

var simpleError = errors.New("simple error")

var _ = Describe("Flusher", func() {
	var (
		mockRepo  *mocks.MockPlanRepo
		ctrl      *gomock.Controller
		chunkSize int
		testData  = []models.Plan{
			testPlan(1),
			testPlan(2),
			testPlan(3),
			testPlan(4),
			testPlan(5),
		}
		fl  flusher.Flusher
		ctx context.Context
	)

	BeforeEach(func() {
		ctx = context.TODO()
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockPlanRepo(ctrl)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	JustBeforeEach(func() {
		fl = flusher.NewFlusher(chunkSize, mockRepo)
	})

	Context("With right chunk size", func() {
		BeforeEach(func() {
			chunkSize = 3
		})
		It("Should works without errors when repo works good", func() {
			mockRepo.EXPECT().AddEntities(gomock.Any(), testData[0:3]).Return(nil).Times(1)
			mockRepo.EXPECT().AddEntities(gomock.Any(), testData[3:5]).Return(nil).Times(1)

			notSaved := fl.Flush(ctx, testData)

			Expect(notSaved).Should(BeNil())
		})

		It("Should return back values at which error occurred", func() {
			mockRepo.EXPECT().AddEntities(gomock.Any(), testData[0:3]).Return(simpleError).Times(1)
			mockRepo.EXPECT().AddEntities(gomock.Any(), testData[3:5]).Return(nil).Times(1)

			notSaved := fl.Flush(ctx, testData)

			Expect(notSaved).Should(Equal(testData[0:3]))
		})

		It("Should return back all values when all add entities in repo failed", func() {
			mockRepo.EXPECT().AddEntities(gomock.Any(), testData[0:3]).Return(simpleError).Times(1)
			mockRepo.EXPECT().AddEntities(gomock.Any(), testData[3:5]).Return(simpleError).Times(1)

			notSaved := fl.Flush(ctx, testData)

			Expect(notSaved).Should(Equal(testData))
		})
	})

	Context("With wrong chunk size", func() {
		BeforeEach(func() {
			chunkSize = -1
		})
		It("Should works with errors", func() {
			fl := flusher.NewFlusher(chunkSize, mockRepo)
			mockRepo.EXPECT().AddEntities(gomock.Any(), gomock.Any()).Times(0)
			notSaved := fl.Flush(ctx, testData)
			Expect(notSaved).To(Equal(testData))
		})
	})
})

func testPlan(id uint64) models.Plan {
	pl := models.NewPlan(id, rand.Uint64(), randStr(), randStr(), time.Now(), randTime())
	return *pl
}

func randStr() string {
	return strconv.FormatUint(rand.Uint64(), 10)
}

func randTime() time.Time {
	durationToAdd := time.Duration(rand.Uint32())
	return time.Now().Add(durationToAdd)
}
