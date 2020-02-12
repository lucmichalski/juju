// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/extensions/v1beta1 (interfaces: ExtensionsV1beta1Interface,IngressInterface)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	v1beta10 "k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockExtensionsV1beta1Interface is a mock of ExtensionsV1beta1Interface interface
type MockExtensionsV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockExtensionsV1beta1InterfaceMockRecorder
}

// MockExtensionsV1beta1InterfaceMockRecorder is the mock recorder for MockExtensionsV1beta1Interface
type MockExtensionsV1beta1InterfaceMockRecorder struct {
	mock *MockExtensionsV1beta1Interface
}

// NewMockExtensionsV1beta1Interface creates a new mock instance
func NewMockExtensionsV1beta1Interface(ctrl *gomock.Controller) *MockExtensionsV1beta1Interface {
	mock := &MockExtensionsV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockExtensionsV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExtensionsV1beta1Interface) EXPECT() *MockExtensionsV1beta1InterfaceMockRecorder {
	return m.recorder
}

// DaemonSets mocks base method
func (m *MockExtensionsV1beta1Interface) DaemonSets(arg0 string) v1beta10.DaemonSetInterface {
	ret := m.ctrl.Call(m, "DaemonSets", arg0)
	ret0, _ := ret[0].(v1beta10.DaemonSetInterface)
	return ret0
}

// DaemonSets indicates an expected call of DaemonSets
func (mr *MockExtensionsV1beta1InterfaceMockRecorder) DaemonSets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DaemonSets", reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).DaemonSets), arg0)
}

// Deployments mocks base method
func (m *MockExtensionsV1beta1Interface) Deployments(arg0 string) v1beta10.DeploymentInterface {
	ret := m.ctrl.Call(m, "Deployments", arg0)
	ret0, _ := ret[0].(v1beta10.DeploymentInterface)
	return ret0
}

// Deployments indicates an expected call of Deployments
func (mr *MockExtensionsV1beta1InterfaceMockRecorder) Deployments(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deployments", reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).Deployments), arg0)
}

// Ingresses mocks base method
func (m *MockExtensionsV1beta1Interface) Ingresses(arg0 string) v1beta10.IngressInterface {
	ret := m.ctrl.Call(m, "Ingresses", arg0)
	ret0, _ := ret[0].(v1beta10.IngressInterface)
	return ret0
}

// Ingresses indicates an expected call of Ingresses
func (mr *MockExtensionsV1beta1InterfaceMockRecorder) Ingresses(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ingresses", reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).Ingresses), arg0)
}

// NetworkPolicies mocks base method
func (m *MockExtensionsV1beta1Interface) NetworkPolicies(arg0 string) v1beta10.NetworkPolicyInterface {
	ret := m.ctrl.Call(m, "NetworkPolicies", arg0)
	ret0, _ := ret[0].(v1beta10.NetworkPolicyInterface)
	return ret0
}

// NetworkPolicies indicates an expected call of NetworkPolicies
func (mr *MockExtensionsV1beta1InterfaceMockRecorder) NetworkPolicies(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkPolicies", reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).NetworkPolicies), arg0)
}

// PodSecurityPolicies mocks base method
func (m *MockExtensionsV1beta1Interface) PodSecurityPolicies() v1beta10.PodSecurityPolicyInterface {
	ret := m.ctrl.Call(m, "PodSecurityPolicies")
	ret0, _ := ret[0].(v1beta10.PodSecurityPolicyInterface)
	return ret0
}

// PodSecurityPolicies indicates an expected call of PodSecurityPolicies
func (mr *MockExtensionsV1beta1InterfaceMockRecorder) PodSecurityPolicies() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PodSecurityPolicies", reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).PodSecurityPolicies))
}

// RESTClient mocks base method
func (m *MockExtensionsV1beta1Interface) RESTClient() rest.Interface {
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient
func (mr *MockExtensionsV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).RESTClient))
}

// ReplicaSets mocks base method
func (m *MockExtensionsV1beta1Interface) ReplicaSets(arg0 string) v1beta10.ReplicaSetInterface {
	ret := m.ctrl.Call(m, "ReplicaSets", arg0)
	ret0, _ := ret[0].(v1beta10.ReplicaSetInterface)
	return ret0
}

// ReplicaSets indicates an expected call of ReplicaSets
func (mr *MockExtensionsV1beta1InterfaceMockRecorder) ReplicaSets(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReplicaSets", reflect.TypeOf((*MockExtensionsV1beta1Interface)(nil).ReplicaSets), arg0)
}

// MockIngressInterface is a mock of IngressInterface interface
type MockIngressInterface struct {
	ctrl     *gomock.Controller
	recorder *MockIngressInterfaceMockRecorder
}

// MockIngressInterfaceMockRecorder is the mock recorder for MockIngressInterface
type MockIngressInterfaceMockRecorder struct {
	mock *MockIngressInterface
}

// NewMockIngressInterface creates a new mock instance
func NewMockIngressInterface(ctrl *gomock.Controller) *MockIngressInterface {
	mock := &MockIngressInterface{ctrl: ctrl}
	mock.recorder = &MockIngressInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIngressInterface) EXPECT() *MockIngressInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockIngressInterface) Create(arg0 *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockIngressInterfaceMockRecorder) Create(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIngressInterface)(nil).Create), arg0)
}

// Delete mocks base method
func (m *MockIngressInterface) Delete(arg0 string, arg1 *v1.DeleteOptions) error {
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockIngressInterfaceMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockIngressInterface)(nil).Delete), arg0, arg1)
}

// DeleteCollection mocks base method
func (m *MockIngressInterface) DeleteCollection(arg0 *v1.DeleteOptions, arg1 v1.ListOptions) error {
	ret := m.ctrl.Call(m, "DeleteCollection", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCollection indicates an expected call of DeleteCollection
func (mr *MockIngressInterfaceMockRecorder) DeleteCollection(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCollection", reflect.TypeOf((*MockIngressInterface)(nil).DeleteCollection), arg0, arg1)
}

// Get mocks base method
func (m *MockIngressInterface) Get(arg0 string, arg1 v1.GetOptions) (*v1beta1.Ingress, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockIngressInterfaceMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockIngressInterface)(nil).Get), arg0, arg1)
}

// List mocks base method
func (m *MockIngressInterface) List(arg0 v1.ListOptions) (*v1beta1.IngressList, error) {
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(*v1beta1.IngressList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockIngressInterfaceMockRecorder) List(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockIngressInterface)(nil).List), arg0)
}

// Patch mocks base method
func (m *MockIngressInterface) Patch(arg0 string, arg1 types.PatchType, arg2 []byte, arg3 ...string) (*v1beta1.Ingress, error) {
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Patch", varargs...)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Patch indicates an expected call of Patch
func (mr *MockIngressInterfaceMockRecorder) Patch(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockIngressInterface)(nil).Patch), varargs...)
}

// Update mocks base method
func (m *MockIngressInterface) Update(arg0 *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockIngressInterfaceMockRecorder) Update(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockIngressInterface)(nil).Update), arg0)
}

// UpdateStatus mocks base method
func (m *MockIngressInterface) UpdateStatus(arg0 *v1beta1.Ingress) (*v1beta1.Ingress, error) {
	ret := m.ctrl.Call(m, "UpdateStatus", arg0)
	ret0, _ := ret[0].(*v1beta1.Ingress)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStatus indicates an expected call of UpdateStatus
func (mr *MockIngressInterfaceMockRecorder) UpdateStatus(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStatus", reflect.TypeOf((*MockIngressInterface)(nil).UpdateStatus), arg0)
}

// Watch mocks base method
func (m *MockIngressInterface) Watch(arg0 v1.ListOptions) (watch.Interface, error) {
	ret := m.ctrl.Call(m, "Watch", arg0)
	ret0, _ := ret[0].(watch.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockIngressInterfaceMockRecorder) Watch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockIngressInterface)(nil).Watch), arg0)
}
