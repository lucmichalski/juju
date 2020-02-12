// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/dynamic (interfaces: Interface,ResourceInterface,NamespaceableResourceInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	dynamic "k8s.io/client-go/dynamic"
	reflect "reflect"
)

// MockDynamicInterface is a mock of Interface interface
type MockDynamicInterface struct {
	ctrl     *gomock.Controller
	recorder *MockDynamicInterfaceMockRecorder
}

// MockDynamicInterfaceMockRecorder is the mock recorder for MockDynamicInterface
type MockDynamicInterfaceMockRecorder struct {
	mock *MockDynamicInterface
}

// NewMockDynamicInterface creates a new mock instance
func NewMockDynamicInterface(ctrl *gomock.Controller) *MockDynamicInterface {
	mock := &MockDynamicInterface{ctrl: ctrl}
	mock.recorder = &MockDynamicInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDynamicInterface) EXPECT() *MockDynamicInterfaceMockRecorder {
	return m.recorder
}

// Resource mocks base method
func (m *MockDynamicInterface) Resource(arg0 schema.GroupVersionResource) dynamic.NamespaceableResourceInterface {
	ret := m.ctrl.Call(m, "Resource", arg0)
	ret0, _ := ret[0].(dynamic.NamespaceableResourceInterface)
	return ret0
}

// Resource indicates an expected call of Resource
func (mr *MockDynamicInterfaceMockRecorder) Resource(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resource", reflect.TypeOf((*MockDynamicInterface)(nil).Resource), arg0)
}

// MockResourceInterface is a mock of ResourceInterface interface
type MockResourceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockResourceInterfaceMockRecorder
}

// MockResourceInterfaceMockRecorder is the mock recorder for MockResourceInterface
type MockResourceInterfaceMockRecorder struct {
	mock *MockResourceInterface
}

// NewMockResourceInterface creates a new mock instance
func NewMockResourceInterface(ctrl *gomock.Controller) *MockResourceInterface {
	mock := &MockResourceInterface{ctrl: ctrl}
	mock.recorder = &MockResourceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourceInterface) EXPECT() *MockResourceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockResourceInterface) Create(arg0 *unstructured.Unstructured, arg1 v1.CreateOptions, arg2 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockResourceInterfaceMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockResourceInterface)(nil).Create), varargs...)
}

// Delete mocks base method
func (m *MockResourceInterface) Delete(arg0 string, arg1 *v1.DeleteOptions, arg2 ...string) error {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockResourceInterfaceMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResourceInterface)(nil).Delete), varargs...)
}

// DeleteCollection mocks base method
func (m *MockResourceInterface) DeleteCollection(arg0 *v1.DeleteOptions, arg1 v1.ListOptions) error {
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockResourceInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockResourceInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockResourceInterface) Get(arg0 string, arg1 v1.GetOptions, arg2 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockResourceInterfaceMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResourceInterface)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockResourceInterface) List(arg0 v1.ListOptions) (*unstructured.UnstructuredList, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockResourceInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockResourceInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 v1.PatchOptions, arg4 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockResourceInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockResourceInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockResourceInterface) Update(arg0 *unstructured.Unstructured, arg1 v1.UpdateOptions, arg2 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockResourceInterfaceMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResourceInterface)(nil).Update), varargs...)
}

// UpdateStatus mocks base method
func (m *MockResourceInterface) UpdateStatus(arg0 *unstructured.Unstructured, arg1 v1.UpdateOptions) (*unstructured.Unstructured, error) {
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockResourceInterfaceMockRecorder) UpdateStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockResourceInterface)(nil).UpdateStatus), arg0, arg1)
}

// Watch mocks base method
func (m *MockResourceInterface) Watch(arg0 v1.ListOptions) (watch.Interface, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockResourceInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockResourceInterface)(nil).Watch), arg0)
}

// MockNamespaceableResourceInterface is a mock of NamespaceableResourceInterface interface
type MockNamespaceableResourceInterface struct {
	ctrl     *gomock.Controller
	recorder *MockNamespaceableResourceInterfaceMockRecorder
}

// MockNamespaceableResourceInterfaceMockRecorder is the mock recorder for MockNamespaceableResourceInterface
type MockNamespaceableResourceInterfaceMockRecorder struct {
	mock *MockNamespaceableResourceInterface
}

// NewMockNamespaceableResourceInterface creates a new mock instance
func NewMockNamespaceableResourceInterface(ctrl *gomock.Controller) *MockNamespaceableResourceInterface {
	mock := &MockNamespaceableResourceInterface{ctrl: ctrl}
	mock.recorder = &MockNamespaceableResourceInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNamespaceableResourceInterface) EXPECT() *MockNamespaceableResourceInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockNamespaceableResourceInterface) Create(arg0 *unstructured.Unstructured, arg1 v1.CreateOptions, arg2 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Create), varargs...)
}

// Delete mocks base method
func (m *MockNamespaceableResourceInterface) Delete(arg0 string, arg1 *v1.DeleteOptions, arg2 ...string) error {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Delete), varargs...)
}

// DeleteCollection mocks base method
func (m *MockNamespaceableResourceInterface) DeleteCollection(arg0 *v1.DeleteOptions, arg1 v1.ListOptions) error {
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockNamespaceableResourceInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockNamespaceableResourceInterface) Get(arg0 string, arg1 v1.GetOptions, arg2 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Get), varargs...)
}

// List mocks base method
func (m *MockNamespaceableResourceInterface) List(arg0 v1.ListOptions) (*unstructured.UnstructuredList, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*unstructured.UnstructuredList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockNamespaceableResourceInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).List), arg0)
}

// Namespace mocks base method
func (m *MockNamespaceableResourceInterface) Namespace(arg0 string) dynamic.ResourceInterface {
	ret := m.ctrl.Call(m, "Namespace", arg0)
	ret0, _ := ret[0].(dynamic.ResourceInterface)
	return ret0
}

// Namespace indicates an expected call of Namespace
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Namespace(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Namespace", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Namespace), arg0)
}

// Patch mocks base method
func (m *MockNamespaceableResourceInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 v1.PatchOptions, arg4 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1, arg2, arg3}
	for _, a := range arg4 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Patch(arg0, arg1, arg2, arg3 interface{}, arg4 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2, arg3}, arg4...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockNamespaceableResourceInterface) Update(arg0 *unstructured.Unstructured, arg1 v1.UpdateOptions, arg2 ...string) (*unstructured.Unstructured, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Update", varargs...)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Update(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Update), varargs...)
}

// UpdateStatus mocks base method
func (m *MockNamespaceableResourceInterface) UpdateStatus(arg0 *unstructured.Unstructured, arg1 v1.UpdateOptions) (*unstructured.Unstructured, error) {
	ret := m.ctrl.Call(m, "UpdateStatus", arg0, arg1)
	ret0, _ := ret[0].(*unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockNamespaceableResourceInterfaceMockRecorder) UpdateStatus(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).UpdateStatus), arg0, arg1)
}

// Watch mocks base method
func (m *MockNamespaceableResourceInterface) Watch(arg0 v1.ListOptions) (watch.Interface, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockNamespaceableResourceInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockNamespaceableResourceInterface)(nil).Watch), arg0)
}
