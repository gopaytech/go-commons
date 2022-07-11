// Code generated by mockery v2.13.1. DO NOT EDIT.

package gitlab_mock

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "github.com/xanzy/go-gitlab"
)

// GroupVariables is an autogenerated mock type for the GroupVariables type
type GroupVariables struct {
	mock.Mock
}

// GetVariable provides a mock function with given fields: gid, key
func (_m *GroupVariables) GetVariable(gid interface{}, key string) (*gitlab.GroupVariable, error) {
	ret := _m.Called(gid, key)

	var r0 *gitlab.GroupVariable
	if rf, ok := ret.Get(0).(func(interface{}, string) *gitlab.GroupVariable); ok {
		r0 = rf(gid, key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.GroupVariable)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, string) error); ok {
		r1 = rf(gid, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGroupVariables interface {
	mock.TestingT
	Cleanup(func())
}

// NewGroupVariables creates a new instance of GroupVariables. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGroupVariables(t mockConstructorTestingTNewGroupVariables) *GroupVariables {
	mock := &GroupVariables{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
