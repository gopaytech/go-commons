// Code generated by mockery v2.9.4. DO NOT EDIT.

package tmpl_mock

import (
	fs "io/fs"

	mock "github.com/stretchr/testify/mock"
)

// FileFilter is an autogenerated mock type for the FileFilter type
type FileFilter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: path, info
func (_m *FileFilter) Execute(path string, info fs.FileInfo) bool {
	ret := _m.Called(path, info)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, fs.FileInfo) bool); ok {
		r0 = rf(path, info)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
