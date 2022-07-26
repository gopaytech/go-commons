// Code generated by mockery v2.14.0. DO NOT EDIT.

package gitlab_mock

import (
	gitlab "github.com/gopaytech/go-commons/pkg/gitlab"
	go_gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// ProtectedBranch is an autogenerated mock type for the ProtectedBranch type
type ProtectedBranch struct {
	mock.Mock
}

// CreateProtectedBranch provides a mock function with given fields: projectID, branchName
func (_m *ProtectedBranch) CreateProtectedBranch(projectID gitlab.NameOrId, branchName string) (*go_gitlab.ProtectedBranch, *go_gitlab.Response, error) {
	ret := _m.Called(projectID, branchName)

	var r0 *go_gitlab.ProtectedBranch
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId, string) *go_gitlab.ProtectedBranch); ok {
		r0 = rf(projectID, branchName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_gitlab.ProtectedBranch)
		}
	}

	var r1 *go_gitlab.Response
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId, string) *go_gitlab.Response); ok {
		r1 = rf(projectID, branchName)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*go_gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(gitlab.NameOrId, string) error); ok {
		r2 = rf(projectID, branchName)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type mockConstructorTestingTNewProtectedBranch interface {
	mock.TestingT
	Cleanup(func())
}

// NewProtectedBranch creates a new instance of ProtectedBranch. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProtectedBranch(t mockConstructorTestingTNewProtectedBranch) *ProtectedBranch {
	mock := &ProtectedBranch{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
