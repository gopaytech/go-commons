// Code generated by mockery v2.9.4. DO NOT EDIT.

package gitlab_mock

import (
	gitlab "github.com/gopaytech/go-commons/pkg/gitlab"
	go_gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// Group is an autogenerated mock type for the Group type
type Group struct {
	mock.Mock
}

// CreateGroup provides a mock function with given fields: name, parent, visibility
func (_m *Group) CreateGroup(name string, parent gitlab.NameOrId, visibility go_gitlab.VisibilityValue) (*go_gitlab.Group, *go_gitlab.Response, error) {
	ret := _m.Called(name, parent, visibility)

	var r0 *go_gitlab.Group
	if rf, ok := ret.Get(0).(func(string, gitlab.NameOrId, go_gitlab.VisibilityValue) *go_gitlab.Group); ok {
		r0 = rf(name, parent, visibility)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_gitlab.Group)
		}
	}

	var r1 *go_gitlab.Response
	if rf, ok := ret.Get(1).(func(string, gitlab.NameOrId, go_gitlab.VisibilityValue) *go_gitlab.Response); ok {
		r1 = rf(name, parent, visibility)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*go_gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, gitlab.NameOrId, go_gitlab.VisibilityValue) error); ok {
		r2 = rf(name, parent, visibility)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetGroup provides a mock function with given fields: id
func (_m *Group) GetGroup(id gitlab.NameOrId) (*go_gitlab.Group, error) {
	ret := _m.Called(id)

	var r0 *go_gitlab.Group
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId) *go_gitlab.Group); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*go_gitlab.Group)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListAllProjects provides a mock function with given fields: id
func (_m *Group) ListAllProjects(id gitlab.NameOrId) (<-chan go_gitlab.Project, error) {
	ret := _m.Called(id)

	var r0 <-chan go_gitlab.Project
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId) <-chan go_gitlab.Project); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan go_gitlab.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(gitlab.NameOrId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}