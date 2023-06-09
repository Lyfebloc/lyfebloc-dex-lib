// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Lyfebloc/lyfebloc-dex-lib/pkg/source/metavault (interfaces: IUSDMReader)

// Package metavault is a generated GoMock package.
package metavault

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIUSDMReader is a mock of IUSDMReader interface.
type MockIUSDMReader struct {
	ctrl     *gomock.Controller
	recorder *MockIUSDMReaderMockRecorder
}

// MockIUSDMReaderMockRecorder is the mock recorder for MockIUSDMReader.
type MockIUSDMReaderMockRecorder struct {
	mock *MockIUSDMReader
}

// NewMockIUSDMReader creates a new mock instance.
func NewMockIUSDMReader(ctrl *gomock.Controller) *MockIUSDMReader {
	mock := &MockIUSDMReader{ctrl: ctrl}
	mock.recorder = &MockIUSDMReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUSDMReader) EXPECT() *MockIUSDMReaderMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockIUSDMReader) Read(arg0 context.Context, arg1 string) (*USDM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0, arg1)
	ret0, _ := ret[0].(*USDM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockIUSDMReaderMockRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockIUSDMReader)(nil).Read), arg0, arg1)
}
