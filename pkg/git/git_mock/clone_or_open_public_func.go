// Code generated by mockery 2.9.0. DO NOT EDIT.

package git_mock

import (
	git "github.com/gopaytech/go-commons/pkg/git"
	mock "github.com/stretchr/testify/mock"
)

// CloneOrOpenPublicFunc is an autogenerated mock type for the CloneOrOpenPublicFunc type
type CloneOrOpenPublicFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: repositoryUrl, destination
func (_m *CloneOrOpenPublicFunc) Execute(repositoryUrl string, destination string) (git.Repository, error) {
	ret := _m.Called(repositoryUrl, destination)

	var r0 git.Repository
	if rf, ok := ret.Get(0).(func(string, string) git.Repository); ok {
		r0 = rf(repositoryUrl, destination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(git.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(repositoryUrl, destination)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}