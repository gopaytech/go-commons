// Code generated by mockery v2.9.4. DO NOT EDIT.

package gitlab_mock

import (
	mock "github.com/stretchr/testify/mock"
	gitlab "github.com/xanzy/go-gitlab"

	pkggitlab "github.com/gopaytech/go-commons/pkg/gitlab"
)

// Project is an autogenerated mock type for the Project type
type Project struct {
	mock.Mock
}

// CreateProject provides a mock function with given fields: name, parentID, visibility
func (_m *Project) CreateProject(name string, parentID int, visibility gitlab.VisibilityValue) (*gitlab.Project, error) {
	ret := _m.Called(name, parentID, visibility)

	var r0 *gitlab.Project
	if rf, ok := ret.Get(0).(func(string, int, gitlab.VisibilityValue) *gitlab.Project); ok {
		r0 = rf(name, parentID, visibility)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, int, gitlab.VisibilityValue) error); ok {
		r1 = rf(name, parentID, visibility)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: id
func (_m *Project) Get(id pkggitlab.NameOrId) (*gitlab.Project, error) {
	ret := _m.Called(id)

	var r0 *gitlab.Project
	if rf, ok := ret.Get(0).(func(pkggitlab.NameOrId) *gitlab.Project); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pkggitlab.NameOrId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultBranch provides a mock function with given fields: id
func (_m *Project) GetDefaultBranch(id pkggitlab.NameOrId) (*gitlab.Branch, error) {
	ret := _m.Called(id)

	var r0 *gitlab.Branch
	if rf, ok := ret.Get(0).(func(pkggitlab.NameOrId) *gitlab.Branch); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Branch)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(pkggitlab.NameOrId) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
