package gitlab

import gl "github.com/xanzy/go-gitlab"

type Search interface {
	ProjectsByGroup(parent NameOrId, name string) ([]*gl.Project, *gl.Response, error)
	ProjectsByName(name string) ([]*gl.Project, *gl.Response, error)
}

type search struct {
	client *gl.Client
}

func (s *search) ProjectsByName(name string) ([]*gl.Project, *gl.Response, error) {
	opts := &gl.SearchOptions{}
	return s.client.Search.Projects(name, opts)
}

func (s *search) ProjectsByGroup(parent NameOrId, name string) ([]*gl.Project, *gl.Response, error) {
	opts := &gl.SearchOptions{}
	return s.client.Search.ProjectsByGroup(parent.ID, name, opts)
}

func NewSearch(client *gl.Client) Search {
	return &search{client: client}
}
