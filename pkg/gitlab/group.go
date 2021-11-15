package gitlab

import (
	"sync"

	gl "github.com/xanzy/go-gitlab"
)

type Group interface {
	ListAllProjects(id NameOrId) (<-chan gl.Project, error)
	GetGroup(id NameOrId) (*gl.Group, error)
}

type group struct {
	client *gl.Client
}

func (g *group) GetGroup(id NameOrId) (*gl.Group, error) {
	getGroupOptions := &gl.GetGroupOptions{}
	result, _, err := g.client.Groups.GetGroup(id.Get(), getGroupOptions)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func NewGroup(client *gl.Client) Group {
	return &group{client: client}
}

func (g *group) ListAllProjects(id NameOrId) (<-chan gl.Project, error) {
	projects := make(chan gl.Project)

	listOpts := gl.ListOptions{
		Page:    0,
		PerPage: 25,
	}

	opts := &gl.ListGroupProjectsOptions{IncludeSubgroups: gl.Bool(true), ListOptions: listOpts}
	result, response, err := g.client.Groups.ListGroupProjects(id.Get(), opts)
	if err != nil {
		return projects, err
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for _, p := range result {
			if p != nil {
				projects <- *p
			}
		}
		wg.Done()
	}()

	for i := response.CurrentPage; i <= response.TotalPages; i++ {
		wg.Add(1)
		go func(page int) {
			lOps := gl.ListOptions{
				Page:    page,
				PerPage: 25,
			}

			iOps := &gl.ListGroupProjectsOptions{IncludeSubgroups: gl.Bool(true), ListOptions: lOps}
			iResult, _, ierr := g.client.Groups.ListGroupProjects(id.Get(), iOps)
			if ierr == nil && iResult != nil {
				for _, p := range iResult {
					if p != nil {
						projects <- *p
					}
				}
			}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(projects)
	}()

	return projects, nil
}
