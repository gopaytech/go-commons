package gitlab

import (
	"fmt"

	gl "github.com/xanzy/go-gitlab"
)

type Project interface {
	Get(id NameOrId) (*gl.Project, error)
	GetDefaultBranch(id NameOrId) (*gl.Branch, error)
}

type project struct {
	client *gl.Client
}

func (p *project) GetDefaultBranch(id NameOrId) (*gl.Branch, error) {
	branches, _, err := p.client.Branches.ListBranches(id.Get(), &gl.ListBranchesOptions{})
	if err != nil {
		return nil, err
	}

	for _, branch := range branches {
		if branch.Default {
			return branch, nil
		}
	}

	return nil, fmt.Errorf("default branch for project %v is not found", id.Get())
}

func (p *project) Get(id NameOrId) (*gl.Project, error) {
	project, _, err := p.client.Projects.GetProject(id.Get(), &gl.GetProjectOptions{})
	if err != nil {
		return nil, err
	}
	return project, err
}

func NewProject(client *gl.Client) Project {
	return &project{client: client}
}
