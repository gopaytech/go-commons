package deb

import (
	"bytes"
	"errors"
	"testing"

	"github.com/gopaytech/go-commons/pkg/file"
	"github.com/goreleaser/nfpm"
	"github.com/stretchr/testify/assert"
)

type builderTestContext struct {
	builder      Builder
	nfpmPackager *NFPMPackagerMock
	fileUtil     *file.FileMock
}

func (context *builderTestContext) setUp(t *testing.T) {
	context.nfpmPackager = &NFPMPackagerMock{}
	context.fileUtil = &file.FileMock{}
	context.builder = &builder{
		packager:                context.nfpmPackager,
		saveFunc:                context.fileUtil.Save,
		getFilesInDirectoryFunc: context.fileUtil.GetFilesInDirectory,
	}
}

func (context *builderTestContext) tearDown() {
}

func TestBuilder_BuildSuccess(t *testing.T) {
	testContext := builderTestContext{}
	testContext.setUp(t)
	defer testContext.tearDown()

	mockBuildSuccess(&testContext)

	config := Config{
		Name:        "test-package",
		Version:     "v1.0.0",
		Destination: "/opt/test-packager",
		PostRemove:  "./postrm.sh",
		Arch:        "amd64",
		Source:      "./package/pkg",
	}
	err := testContext.builder.Build(config, "/tmp/result.deb")

	assert.NoError(t, err)
	testContext.nfpmPackager.AssertExpectations(t)
	testContext.fileUtil.AssertExpectations(t)
}

func TestBuilder_BuildFileDiscoveryError(t *testing.T) {
	testContext := builderTestContext{}
	testContext.setUp(t)
	defer testContext.tearDown()

	mockBuildFileDiscoveryError(&testContext)

	config := Config{
		Name:        "test-package",
		Version:     "v1.0.0",
		Destination: "/opt/test-packager",
		PostRemove:  "./postrm.sh",
		Arch:        "amd64",
		Source:      "./package/pkg",
	}
	err := testContext.builder.Build(config, "/tmp/result.deb")

	assert.Error(t, err)
	testContext.nfpmPackager.AssertExpectations(t)
	testContext.fileUtil.AssertExpectations(t)
}

func TestBuilder_BuildPackageError(t *testing.T) {
	testContext := builderTestContext{}
	testContext.setUp(t)
	defer testContext.tearDown()

	mockBuildPackageError(&testContext)

	config := Config{
		Name:        "test-package",
		Version:     "v1.0.0",
		Destination: "/opt/test-packager",
		PostRemove:  "./postrm.sh",
		Arch:        "amd64",
		Source:      "./package/pkg",
	}
	err := testContext.builder.Build(config, "/tmp/result.deb")

	assert.Error(t, err)
	testContext.nfpmPackager.AssertExpectations(t)
	testContext.fileUtil.AssertExpectations(t)
}

func mockBuildSuccess(testContext *builderTestContext) {
	nfpmInfo := nfpm.Info{
		Name:        "test-package",
		Version:     "v1.0.0",
		Bindir:      "/opt/test-packager",
		Arch:        "amd64",
		Maintainer:  "Gopay-Systems",
		Description: "test-package_v1.0.0",
		Overridables: nfpm.Overridables{
			Scripts: nfpm.Scripts{
				PostRemove: "./postrm.sh",
			},
			Files: map[string]string{
				"./package/pkg/binary": "/opt/test-packager/binary",
			},
		},
	}
	writer := bytes.Buffer{}
	testContext.nfpmPackager.
		On("Package", &nfpmInfo, &writer).
		Return(nil)
	testContext.fileUtil.
		On("GetFilesInDirectory", "./package/pkg").
		Return([]string{"binary"}, nil)
	testContext.fileUtil.
		On("Save", writer.Bytes(), "/tmp/result.deb").
		Return(nil)
}

func mockBuildFileDiscoveryError(testContext *builderTestContext) {
	testContext.fileUtil.
		On("GetFilesInDirectory", "./package/pkg").
		Return([]string{"binary"}, errors.New("unable to access directory"))
}

func mockBuildPackageError(testContext *builderTestContext) {
	nfpmInfo := nfpm.Info{
		Name:        "test-package",
		Version:     "v1.0.0",
		Bindir:      "/opt/test-packager",
		Arch:        "amd64",
		Maintainer:  "Gopay-Systems",
		Description: "test-package_v1.0.0",
		Overridables: nfpm.Overridables{
			Scripts: nfpm.Scripts{
				PostRemove: "./postrm.sh",
			},
			Files: map[string]string{
				"./package/pkg/binary": "/opt/test-packager/binary",
			},
		},
	}
	writer := bytes.Buffer{}
	testContext.nfpmPackager.
		On("Package", &nfpmInfo, &writer).
		Return(errors.New("unable to package deb file"))
	testContext.fileUtil.
		On("GetFilesInDirectory", "./package/pkg").
		Return([]string{"binary"}, nil)
}
