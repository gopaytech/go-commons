package deb

import (
	"bytes"
	"fmt"

	"github.com/gopaytech/go-commons/pkg/file"
	"github.com/goreleaser/nfpm"
	"github.com/goreleaser/nfpm/deb"
)

type Builder interface {
	Build(config Config, filename string) (err error)
}

type builder struct {
	packager                nfpm.Packager
	getFilesInDirectoryFunc file.GetFilesInDirectoryFunc
	saveFunc                file.SaveFunc
}

func (b *builder) Build(config Config, filename string) (err error) {
	mappedFiles, err := b.mapFiles(config.Source, config.Destination)
	if err != nil {
		return
	}
	nfpmConfig := config.ConvertToNFPMConfig()
	nfpmConfig.Files = mappedFiles
	contentBuffer := bytes.Buffer{}
	err = b.packager.Package(nfpmConfig, &contentBuffer)
	if err != nil {
		return
	}
	err = b.saveFunc(contentBuffer.Bytes(), filename)
	return
}

func (b *builder) mapFiles(source string, target string) (mappedFiles map[string]string, err error) {
	files, err := b.getFilesInDirectoryFunc(source)
	if err != nil {
		return
	}
	mappedFiles = map[string]string{}
	for _, _file := range files {
		sourceFiles := fmt.Sprintf("%s/%s", source, _file)
		targetFiles := fmt.Sprintf("%s/%s", target, _file)
		mappedFiles[sourceFiles] = targetFiles
	}
	return
}

func NewBuilder() Builder {
	return &builder{
		packager:                &deb.Deb{},
		saveFunc:                file.Save,
		getFilesInDirectoryFunc: file.GetFilesInDirectory,
	}
}
