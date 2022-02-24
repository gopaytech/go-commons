// Code generated by mockery 2.9.0. DO NOT EDIT.

package git_mock

import (
	git "github.com/go-git/go-git/v5"
	mock "github.com/stretchr/testify/mock"

	pkggit "github.com/gopaytech/go-commons/pkg/git"
)

// CloneWithOptionFunc is an autogenerated mock type for the CloneWithOptionFunc type
type CloneWithOptionFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: destination, option
func (_m *CloneWithOptionFunc) Execute(destination string, option *git.CloneOptions) (pkggit.Repository, error) {
	ret := _m.Called(destination, option)

	var r0 pkggit.Repository
	if rf, ok := ret.Get(0).(func(string, *git.CloneOptions) pkggit.Repository); ok {
		r0 = rf(destination, option)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(pkggit.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *git.CloneOptions) error); ok {
		r1 = rf(destination, option)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}