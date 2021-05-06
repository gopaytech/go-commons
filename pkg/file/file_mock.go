package file

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type FileMock struct {
	mock.Mock
}

func (f *FileMock) JoinPaths(paths ...string) (path string) {
	arguments := f.Called(paths)
	return arguments.String(0)
}

func (f *FileMock) MoveDirectoryContents(sourceDirectory string, targetDirectory string) (err error) {
	arguments := f.Called(sourceDirectory, targetDirectory)
	return arguments.Error(0)
}

func (f *FileMock) GetFoldersInDirectory(directory string) (directories []string, err error) {
	arguments := f.Called(directory)
	return arguments.Get(0).([]string), arguments.Error(1)
}

func (f *FileMock) CreateFolder(folder string) (err error) {
	arguments := f.Called(folder)
	return arguments.Error(0)
}

func (f *FileMock) Tar(sourceDirectory string, writer io.Writer) (err error) {
	arguments := f.Called(sourceDirectory, writer)
	return arguments.Error(0)
}

func (f *FileMock) Move(sourceFile string, destinationFile string) (err error) {
	arguments := f.Called(sourceFile, destinationFile)
	return arguments.Error(0)
}

func (f *FileMock) Copy(sourceFile string, destinationFile string) (err error) {
	arguments := f.Called(sourceFile, destinationFile)
	return arguments.Error(0)
}

func (f *FileMock) GetBasename(filename string) (basename string) {
	arguments := f.Called(filename)
	return arguments.Get(0).(string)
}

func (f *FileMock) GetFilesInDirectory(directory string) (files []string, err error) {
	arguments := f.Called(directory)
	return arguments.Get(0).([]string), arguments.Error(1)
}

func (f *FileMock) Unzip(filename string, directory string) (err error) {
	arguments := f.Called(filename, directory)
	return arguments.Error(0)
}

func (f *FileMock) Save(content []byte, filename string) (err error) {
	arguments := f.Called(content, filename)
	return arguments.Error(0)
}
