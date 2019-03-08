// Package mock_rule is a generated GoMock package.
package mock_rule

import (
	goquery "github.com/PuerkitoBio/goquery"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRule is a mock of Rule interface
type MockRule struct {
	ctrl     *gomock.Controller
	recorder *MockRuleMockRecorder
}

// MockRuleMockRecorder is the mock recorder for MockRule
type MockRuleMockRecorder struct {
	mock *MockRule
}

// NewMockRule creates a new mock instance
func NewMockRule(ctrl *gomock.Controller) *MockRule {
	mock := &MockRule{ctrl: ctrl}
	mock.recorder = &MockRuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRule) EXPECT() *MockRuleMockRecorder {
	return m.recorder
}

// DataRule mocks base method
func (m *MockRule) DataRule(arg0 *goquery.Document, arg1 func(interface{})) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DataRule", arg0, arg1)
}

// DataRule indicates an expected call of DataRule
func (mr *MockRuleMockRecorder) DataRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DataRule", reflect.TypeOf((*MockRule)(nil).DataRule), arg0, arg1)
}

// ImageRule mocks base method
func (m *MockRule) ImageRule(arg0 *goquery.Document, arg1 func(string)) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ImageRule", arg0, arg1)
}

// ImageRule indicates an expected call of ImageRule
func (mr *MockRuleMockRecorder) ImageRule(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImageRule", reflect.TypeOf((*MockRule)(nil).ImageRule), arg0, arg1)
}

// PageRule mocks base method
func (m *MockRule) PageRule(arg0 int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PageRule", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// PageRule indicates an expected call of PageRule
func (mr *MockRuleMockRecorder) PageRule(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PageRule", reflect.TypeOf((*MockRule)(nil).PageRule), arg0)
}

// UrlRule mocks base method
func (m *MockRule) UrlRule() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UrlRule")
	ret0, _ := ret[0].(string)
	return ret0
}

// UrlRule indicates an expected call of UrlRule
func (mr *MockRuleMockRecorder) UrlRule() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UrlRule", reflect.TypeOf((*MockRule)(nil).UrlRule))
}
