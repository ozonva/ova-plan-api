// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozonva/ova-plan-api/internal/repo (interfaces: PlanRepo)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozonva/ova-plan-api/internal/models"
)

// MockPlanRepo is a mock of PlanRepo interface.
type MockPlanRepo struct {
	ctrl     *gomock.Controller
	recorder *MockPlanRepoMockRecorder
}

// MockPlanRepoMockRecorder is the mock recorder for MockPlanRepo.
type MockPlanRepoMockRecorder struct {
	mock *MockPlanRepo
}

// NewMockPlanRepo creates a new mock instance.
func NewMockPlanRepo(ctrl *gomock.Controller) *MockPlanRepo {
	mock := &MockPlanRepo{ctrl: ctrl}
	mock.recorder = &MockPlanRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanRepo) EXPECT() *MockPlanRepoMockRecorder {
	return m.recorder
}

// AddEntities mocks base method.
func (m *MockPlanRepo) AddEntities(arg0 []models.Plan) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntities", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntities indicates an expected call of AddEntities.
func (mr *MockPlanRepoMockRecorder) AddEntities(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntities", reflect.TypeOf((*MockPlanRepo)(nil).AddEntities), arg0)
}

// DescribeEntity mocks base method.
func (m *MockPlanRepo) DescribeEntity(arg0 uint64) (*models.Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntity", arg0)
	ret0, _ := ret[0].(*models.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntity indicates an expected call of DescribeEntity.
func (mr *MockPlanRepoMockRecorder) DescribeEntity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntity", reflect.TypeOf((*MockPlanRepo)(nil).DescribeEntity), arg0)
}

// ListEntities mocks base method.
func (m *MockPlanRepo) ListEntities(arg0, arg1 uint64) ([]models.Plan, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntities", arg0, arg1)
	ret0, _ := ret[0].([]models.Plan)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntities indicates an expected call of ListEntities.
func (mr *MockPlanRepoMockRecorder) ListEntities(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntities", reflect.TypeOf((*MockPlanRepo)(nil).ListEntities), arg0, arg1)
}