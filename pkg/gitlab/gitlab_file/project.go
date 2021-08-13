package gitlab_file

import (
	"encoding/base64"

	"github.com/gopaytech/go-commons/pkg/gitlab"
	gl "github.com/xanzy/go-gitlab"
)

type ProjectFileOps interface {
	SearchFilesMetadata(id gitlab.NameOrId, ref string, fileNames ...string) (filesFound []gl.File, err error)
	SearchFiles(id gitlab.NameOrId, ref string, fileNames ...string) (filesFound []gl.File, err error)
	SearchFileWithContent(id gitlab.NameOrId, ref string, searchFunc SearchFunc, fileNames ...string) (filesFound []gl.File, err error)
}

type projectFileOps struct {
	client *gl.Client
}

func (p *projectFileOps) SearchFileWithContent(id gitlab.NameOrId, ref string, searchFunc SearchFunc, fileNames ...string) (filesFound []gl.File, err error) {

	contentContains := func(base64Content string, fun SearchFunc) bool {
		plainBytes, ierr := base64.StdEncoding.DecodeString(base64Content)
		if ierr != nil {
			return false
		}

		return fun(plainBytes)
	}

	filesFound = []gl.File{}

	files, err := p.SearchFiles(id, ref, fileNames...)
	if err != nil {
		return
	}

	for _, file := range files {
		if contentContains(file.Content, searchFunc) {
			filesFound = append(filesFound, file)
		}
	}

	return
}

func (p *projectFileOps) SearchFilesMetadata(id gitlab.NameOrId, ref string, fileNames ...string) (filesFound []gl.File, err error) {
	filesFound = []gl.File{}
	if len(ref) == 0 {
		ref = "master"
	}

	for _, fileName := range fileNames {
		ops := &gl.GetFileMetaDataOptions{Ref: gl.String(ref)}
		fileFound, response, ierr := p.client.RepositoryFiles.GetFileMetaData(id.Get(), fileName, ops)
		if response != nil && response.StatusCode >= 300 {
			continue
		} else if ierr != nil {
			err = ierr
			return
		}

		if fileFound != nil {
			filesFound = append(filesFound, *fileFound)
		}
	}

	return
}

func (p *projectFileOps) SearchFiles(id gitlab.NameOrId, ref string, fileNames ...string) (filesFound []gl.File, err error) {
	filesFound = []gl.File{}
	if len(ref) == 0 {
		ref = "master"
	}

	for _, fileName := range fileNames {
		ops := &gl.GetFileOptions{Ref: gl.String(ref)}
		fileFound, response, ierr := p.client.RepositoryFiles.GetFile(id.Get(), fileName, ops)
		if response != nil && response.StatusCode >= 300 {
			continue
		} else if ierr != nil {
			err = ierr
			return
		}

		if fileFound != nil {
			filesFound = append(filesFound, *fileFound)
		}
	}

	return
}

func NewProjectFileOps(client *gl.Client) ProjectFileOps {
	return &projectFileOps{client: client}

}
