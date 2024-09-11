// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/cloud/ecsinterface.go

// Package cloud is a generated GoMock package.
package cloud

import (
	reflect "reflect"

	ecs "github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
)

// MockECSInterface is a mock of ECSInterface interface.
type MockECSInterface struct {
	ctrl     *gomock.Controller
	recorder *MockECSInterfaceMockRecorder
}

// MockECSInterfaceMockRecorder is the mock recorder for MockECSInterface.
type MockECSInterfaceMockRecorder struct {
	mock *MockECSInterface
}

// NewMockECSInterface creates a new mock instance.
func NewMockECSInterface(ctrl *gomock.Controller) *MockECSInterface {
	mock := &MockECSInterface{ctrl: ctrl}
	mock.recorder = &MockECSInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECSInterface) EXPECT() *MockECSInterfaceMockRecorder {
	return m.recorder
}

// CreateDisk mocks base method.
func (m *MockECSInterface) CreateDisk(request *ecs.CreateDiskRequest) (*ecs.CreateDiskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDisk", request)
	ret0, _ := ret[0].(*ecs.CreateDiskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDisk indicates an expected call of CreateDisk.
func (mr *MockECSInterfaceMockRecorder) CreateDisk(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDisk", reflect.TypeOf((*MockECSInterface)(nil).CreateDisk), request)
}

// DeleteDisk mocks base method.
func (m *MockECSInterface) DeleteDisk(request *ecs.DeleteDiskRequest) (*ecs.DeleteDiskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDisk", request)
	ret0, _ := ret[0].(*ecs.DeleteDiskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDisk indicates an expected call of DeleteDisk.
func (mr *MockECSInterfaceMockRecorder) DeleteDisk(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDisk", reflect.TypeOf((*MockECSInterface)(nil).DeleteDisk), request)
}

// DescribeAvailableResource mocks base method.
func (m *MockECSInterface) DescribeAvailableResource(request *ecs.DescribeAvailableResourceRequest) (*ecs.DescribeAvailableResourceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeAvailableResource", request)
	ret0, _ := ret[0].(*ecs.DescribeAvailableResourceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeAvailableResource indicates an expected call of DescribeAvailableResource.
func (mr *MockECSInterfaceMockRecorder) DescribeAvailableResource(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeAvailableResource", reflect.TypeOf((*MockECSInterface)(nil).DescribeAvailableResource), request)
}

// DescribeDisks mocks base method.
func (m *MockECSInterface) DescribeDisks(request *ecs.DescribeDisksRequest) (*ecs.DescribeDisksResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeDisks", request)
	ret0, _ := ret[0].(*ecs.DescribeDisksResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeDisks indicates an expected call of DescribeDisks.
func (mr *MockECSInterfaceMockRecorder) DescribeDisks(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeDisks", reflect.TypeOf((*MockECSInterface)(nil).DescribeDisks), request)
}

// DescribeInstanceTypes mocks base method.
func (m *MockECSInterface) DescribeInstanceTypes(request *ecs.DescribeInstanceTypesRequest) (*ecs.DescribeInstanceTypesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstanceTypes", request)
	ret0, _ := ret[0].(*ecs.DescribeInstanceTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstanceTypes indicates an expected call of DescribeInstanceTypes.
func (mr *MockECSInterfaceMockRecorder) DescribeInstanceTypes(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstanceTypes", reflect.TypeOf((*MockECSInterface)(nil).DescribeInstanceTypes), request)
}

// DescribeInstances mocks base method.
func (m *MockECSInterface) DescribeInstances(request *ecs.DescribeInstancesRequest) (*ecs.DescribeInstancesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeInstances", request)
	ret0, _ := ret[0].(*ecs.DescribeInstancesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeInstances indicates an expected call of DescribeInstances.
func (mr *MockECSInterfaceMockRecorder) DescribeInstances(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeInstances", reflect.TypeOf((*MockECSInterface)(nil).DescribeInstances), request)
}

// DescribeSnapshots mocks base method.
func (m *MockECSInterface) DescribeSnapshots(request *ecs.DescribeSnapshotsRequest) (*ecs.DescribeSnapshotsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeSnapshots", request)
	ret0, _ := ret[0].(*ecs.DescribeSnapshotsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeSnapshots indicates an expected call of DescribeSnapshots.
func (mr *MockECSInterfaceMockRecorder) DescribeSnapshots(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeSnapshots", reflect.TypeOf((*MockECSInterface)(nil).DescribeSnapshots), request)
}

// ResizeDisk mocks base method.
func (m *MockECSInterface) ResizeDisk(request *ecs.ResizeDiskRequest) (*ecs.ResizeDiskResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResizeDisk", request)
	ret0, _ := ret[0].(*ecs.ResizeDiskResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ResizeDisk indicates an expected call of ResizeDisk.
func (mr *MockECSInterfaceMockRecorder) ResizeDisk(request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResizeDisk", reflect.TypeOf((*MockECSInterface)(nil).ResizeDisk), request)
}
