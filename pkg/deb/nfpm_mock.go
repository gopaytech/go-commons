package deb

import (
	"io"

	"github.com/goreleaser/nfpm"
	"github.com/stretchr/testify/mock"
)

type NFPMPackagerMock struct {
	mock.Mock
}

func (m *NFPMPackagerMock) ConventionalFileName(info *nfpm.Info) string {
	arguments := m.Called(info)
	return arguments.String(0)
}

func (m *NFPMPackagerMock) Package(info *nfpm.Info, w io.Writer) error {
	arguments := m.Called(info, w)
	return arguments.Error(0)
}
