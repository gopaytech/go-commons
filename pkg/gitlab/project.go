package gitlab

import (
	"fmt"

	gl "github.com/xanzy/go-gitlab"
)

type Project interface {
	Get(id NameOrId) (*gl.Project, error)
	GetDefaultBranch(id NameOrId) (*gl.Branch, error)
	GetBranchByName(id NameOrId, name string) (*gl.Branch, error)
	CreateProject(name string, parentID int, visibility gl.VisibilityValue, defaultBranch string) (*gl.Project, error)
	CreateMinimalMRApproval(id NameOrId, minimalApproval int) (*gl.Project, error)
}

type project struct {
	client *gl.Client
}

func (p *project) CreateMinimalMRApproval(id NameOrId, minimalApproval int) (*gl.Project, error) {
	project, _, err := p.client.Projects.EditProject(id.Get(), &gl.EditProjectOptions{
		ApprovalsBeforeMerge: &minimalApproval,
	})
	return project, err
}

func (p *project) CreateProject(name string, parentID int, visibility gl.VisibilityValue, defaultBranch string) (*gl.Project, error) {
	initializeWithReadme := true
	project, _, err := p.client.Projects.CreateProject(&gl.CreateProjectOptions{
		Name:                 &name,
		NamespaceID:          &parentID,
		Path:                 nil,
		Visibility:           gl.Visibility(visibility),
		InitializeWithReadme: &initializeWithReadme,
		DefaultBranch:        &defaultBranch,
	})
	return project, err
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

func (p *project) GetBranchByName(id NameOrId, name string) (*gl.Branch, error) {
	branch, _, err := p.client.Branches.GetBranch(id.Get(), name)
	if err != nil {
		return nil, err
	}

	return branch, nil
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
