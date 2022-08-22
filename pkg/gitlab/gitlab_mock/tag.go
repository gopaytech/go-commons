// Code generated by mockery v2.14.0. DO NOT EDIT.

package gitlab_mock

import (
	gitlab "github.com/gopaytech/go-commons/pkg/gitlab"
	go_gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// Tag is an autogenerated mock type for the Tag type
type Tag struct {
	mock.Mock
}

// GetLatestTag provides a mock function with given fields: pid, search
func (_m *Tag) GetLatestTag(pid gitlab.NameOrId, search string) (*go_gitlab.Tag, error) {
	ret := _m.Called(pid, search)

	var r0 *go_gitlab.Tag
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId, string) *go_gitlab.Tag); ok {
		r0 = rf(pid, search)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_gitlab.Tag)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId, string) error); ok {
		r1 = rf(pid, search)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTag interface {
	mock.TestingT
	Cleanup(func())
}

// NewTag creates a new instance of Tag. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTag(t mockConstructorTestingTNewTag) *Tag {
	mock := &Tag{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
