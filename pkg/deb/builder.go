package deb

import (
	"bytes"
	"fmt"
	"github.com/goreleaser/nfpm/v2/files"

	"github.com/gopaytech/go-commons/pkg/file"
	"github.com/goreleaser/nfpm/v2"
	"github.com/goreleaser/nfpm/v2/deb"
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
	mappedContent, err := b.mapContent(config.Source, config.Destination)
	if err != nil {
		return
	}
	nfpmConfig := config.ConvertToNFPMConfig()
	nfpmConfig.Overridables.Contents = mappedContent
	contentBuffer := bytes.Buffer{}
	err = b.packager.Package(nfpmConfig, &contentBuffer)
	if err != nil {
		return
	}
	err = b.saveFunc(contentBuffer.Bytes(), filename)
	return
}

func (b *builder) mapContent(source string, target string) (mappedContent []*files.Content, err error) {
	sourceFiles, err := b.getFilesInDirectoryFunc(source)
	if err != nil {
		return
	}
	mappedContent = []*files.Content{}
	for _, _file := range sourceFiles {
		sourceFile := fmt.Sprintf("%s/%s", source, _file)
		targetFile := fmt.Sprintf("%s/%s", target, _file)
		content := &files.Content{
			Source:      sourceFile,
			Destination: targetFile,
		}
		mappedContent = append(mappedContent, content)
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
