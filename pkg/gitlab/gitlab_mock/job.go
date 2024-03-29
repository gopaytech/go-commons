// Code generated by mockery v2.28.1. DO NOT EDIT.

package gitlab_mock

import (
	gitlab "github.com/gopaytech/go-commons/pkg/gitlab"
	go_gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// Job is an autogenerated mock type for the Job type
type Job struct {
	mock.Mock
}

// GetByPipelineID provides a mock function with given fields: pid, pipelineID
func (_m *Job) GetByPipelineID(pid gitlab.NameOrId, pipelineID int) ([]*go_gitlab.Job, error) {
	ret := _m.Called(pid, pipelineID)

	var r0 []*go_gitlab.Job
	var r1 error
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId, int) ([]*go_gitlab.Job, error)); ok {
		return rf(pid, pipelineID)
	}
	if rf, ok := ret.Get(0).(func(gitlab.NameOrId, int) []*go_gitlab.Job); ok {
		r0 = rf(pid, pipelineID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*go_gitlab.Job)
		}
	}

	if rf, ok := ret.Get(1).(func(gitlab.NameOrId, int) error); ok {
		r1 = rf(pid, pipelineID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewJob interface {
	mock.TestingT
	Cleanup(func())
}

// NewJob creates a new instance of Job. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJob(t mockConstructorTestingTNewJob) *Job {
	mock := &Job{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
