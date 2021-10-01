//go:build integration
// +build integration

package deb

import (
	"fmt"
	"testing"

	"github.com/gopaytech/go-commons/pkg/file"
	"github.com/gopaytech/go-commons/pkg/strings"
	"github.com/goreleaser/nfpm/deb"
	"github.com/stretchr/testify/assert"
)

type builderIntegrationContext struct {
	builder Builder
}

func (context *builderIntegrationContext) setUp(t *testing.T) {
	context.builder = &builder{
		packager:                &deb.Deb{},
		saveFunc:                file.Save,
		getFilesInDirectoryFunc: file.GetFilesInDirectory,
	}
}

func (context *builderIntegrationContext) tearDown() {
}

func TestBuilder_Build(t *testing.T) {
	context := builderIntegrationContext{}
	context.setUp(t)
	defer context.tearDown()

	buildConfig := Config{
		Name:        "test-build",
		Version:     "1.0.0-commit-id",
		Destination: "/opt/test-build",
		PostRemove:  "./builder.go",
		Arch:        "amd64",
		Source:      ".",
	}

	filePath := fmt.Sprintf("/tmp/build-result-%s.deb", strings.RandomAlphanumeric(8))
	err := context.builder.Build(buildConfig, filePath)
	assert.NoError(t, err)
	assert.FileExists(t, filePath)
}
