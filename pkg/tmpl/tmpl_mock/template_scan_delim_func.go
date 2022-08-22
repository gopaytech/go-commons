// Code generated by mockery v2.14.0. DO NOT EDIT.

package tmpl_mock

import (
	tmpl "github.com/gopaytech/go-commons/pkg/tmpl"
	mock "github.com/stretchr/testify/mock"
)

// TemplateScanDelimFunc is an autogenerated mock type for the TemplateScanDelimFunc type
type TemplateScanDelimFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: scanPath, tmplExt, startDelims, endDelims
func (_m *TemplateScanDelimFunc) Execute(scanPath string, tmplExt string, startDelims string, endDelims string) (tmpl.ScanResult, error) {
	ret := _m.Called(scanPath, tmplExt, startDelims, endDelims)

	var r0 tmpl.ScanResult
	if rf, ok := ret.Get(0).(func(string, string, string, string) tmpl.ScanResult); ok {
		r0 = rf(scanPath, tmplExt, startDelims, endDelims)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(tmpl.ScanResult)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, string, string) error); ok {
		r1 = rf(scanPath, tmplExt, startDelims, endDelims)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTemplateScanDelimFunc interface {
	mock.TestingT
	Cleanup(func())
}

// NewTemplateScanDelimFunc creates a new instance of TemplateScanDelimFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTemplateScanDelimFunc(t mockConstructorTestingTNewTemplateScanDelimFunc) *TemplateScanDelimFunc {
	mock := &TemplateScanDelimFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
