package gitlab

import (
	"github.com/gopaytech/go-commons/pkg/ptr"
	gl "github.com/xanzy/go-gitlab"
)

type RepositoryFile interface {
	GetFileByPath(pid NameOrId, path, ref string) (*gl.File, error)
	GetRawFileByPath(pid NameOrId, path, ref string) ([]byte, error)
}

type repositoryFile struct {
	client *gl.Client
}

func (f *repositoryFile) GetFileByPath(pid NameOrId, path, ref string) (*gl.File, error) {
	file, _, err := f.client.RepositoryFiles.GetFile(pid.ID, path, &gl.GetFileOptions{
		Ref: ptr.String(ref),
	})
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *repositoryFile) GetRawFileByPath(pid NameOrId, path, ref string) ([]byte, error) {
	fileBytes, _, err := f.client.RepositoryFiles.GetRawFile(pid.ID, path, &gl.GetRawFileOptions{
		Ref: ptr.String(ref),
	})
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func NewRepositoryFile(client *gl.Client) RepositoryFile {
	return &repositoryFile{client: client}
}
