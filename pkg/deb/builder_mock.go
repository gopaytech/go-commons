package deb

import "github.com/stretchr/testify/mock"

type BuilderMock struct {
	mock.Mock
}

func (b *BuilderMock) Build(config Config, filename string) (err error) {
	arguments := b.Called(config, filename)
	return arguments.Error(0)
}
