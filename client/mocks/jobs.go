// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-k8s/client (interfaces: JobsClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/batch/v1"
	v10 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// MockJobsClient is a mock of JobsClient interface.
type MockJobsClient struct {
	ctrl     *gomock.Controller
	recorder *MockJobsClientMockRecorder
}

// MockJobsClientMockRecorder is the mock recorder for MockJobsClient.
type MockJobsClientMockRecorder struct {
	mock *MockJobsClient
}

// NewMockJobsClient creates a new mock instance.
func NewMockJobsClient(ctrl *gomock.Controller) *MockJobsClient {
	mock := &MockJobsClient{ctrl: ctrl}
	mock.recorder = &MockJobsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJobsClient) EXPECT() *MockJobsClientMockRecorder {
	return m.recorder
}

// List mocks base method.
func (m *MockJobsClient) List(arg0 context.Context, arg1 v10.ListOptions) (*v1.JobList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*v1.JobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockJobsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockJobsClient)(nil).List), arg0, arg1)
}