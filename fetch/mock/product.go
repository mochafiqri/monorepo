// Code generated by MockGen. DO NOT EDIT.
// Source: ./commons/interfaces/product.go

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	dtos "github.com/mochafiqri/monorepo/fetch/commons/dtos"
	entities "github.com/mochafiqri/monorepo/fetch/commons/entities"
)

// MockProductRepo is a mock of ProductRepo interface.
type MockProductRepo struct {
	ctrl     *gomock.Controller
	recorder *MockProductRepoMockRecorder
}

// MockProductRepoMockRecorder is the mock recorder for MockProductRepo.
type MockProductRepoMockRecorder struct {
	mock *MockProductRepo
}

// NewMockProductRepo creates a new mock instance.
func NewMockProductRepo(ctrl *gomock.Controller) *MockProductRepo {
	mock := &MockProductRepo{ctrl: ctrl}
	mock.recorder = &MockProductRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductRepo) EXPECT() *MockProductRepoMockRecorder {
	return m.recorder
}

// GetProduct mocks base method.
func (m *MockProductRepo) GetProduct() ([]entities.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct")
	ret0, _ := ret[0].([]entities.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductRepoMockRecorder) GetProduct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductRepo)(nil).GetProduct))
}

// MockProductDomain is a mock of ProductDomain interface.
type MockProductDomain struct {
	ctrl     *gomock.Controller
	recorder *MockProductDomainMockRecorder
}

// MockProductDomainMockRecorder is the mock recorder for MockProductDomain.
type MockProductDomainMockRecorder struct {
	mock *MockProductDomain
}

// NewMockProductDomain creates a new mock instance.
func NewMockProductDomain(ctrl *gomock.Controller) *MockProductDomain {
	mock := &MockProductDomain{ctrl: ctrl}
	mock.recorder = &MockProductDomainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductDomain) EXPECT() *MockProductDomainMockRecorder {
	return m.recorder
}

// GetProduct mocks base method.
func (m *MockProductDomain) GetProduct() dtos.StandardResponseReq {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct")
	ret0, _ := ret[0].(dtos.StandardResponseReq)
	return ret0
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockProductDomainMockRecorder) GetProduct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockProductDomain)(nil).GetProduct))
}

// GetProductRecommended mocks base method.
func (m *MockProductDomain) GetProductRecommended() dtos.StandardResponseReq {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductRecommended")
	ret0, _ := ret[0].(dtos.StandardResponseReq)
	return ret0
}

// GetProductRecommended indicates an expected call of GetProductRecommended.
func (mr *MockProductDomainMockRecorder) GetProductRecommended() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductRecommended", reflect.TypeOf((*MockProductDomain)(nil).GetProductRecommended))
}
