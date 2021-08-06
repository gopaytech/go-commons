package gitlab

import (
	"github.com/gopaytech/go-commons/pkg/strings"
	gl "github.com/xanzy/go-gitlab"
)

type Pipeline interface {
	GetBySHA(id NameOrId, sha string) ([]gl.PipelineInfo, error)
	GetBySHAAndRef(id NameOrId, sha string, ref string) ([]gl.PipelineInfo, error)
	GetBySHAOnDefault(id NameOrId, sha string) ([]gl.PipelineInfo, error)
}

type pipeline struct {
	client  *gl.Client
	project Project
}

func (p *pipeline) GetBySHA(id NameOrId, sha string) ([]gl.PipelineInfo, error) {
	return p.GetBySHAAndRef(id, sha, "")
}

func (p *pipeline) GetBySHAOnDefault(id NameOrId, sha string) ([]gl.PipelineInfo, error) {
	branch, err := p.project.GetDefaultBranch(id)
	if err != nil {
		return nil, err
	}
	return p.GetBySHAAndRef(id, sha, branch.Name)

}

func (p *pipeline) GetBySHAAndRef(id NameOrId, sha string, ref string) ([]gl.PipelineInfo, error) {
	refP := new(string)

	if !strings.IsStringEmpty(ref) {
		*refP = ref
	}

	var pipelines []gl.PipelineInfo
	pipelinesInfos, _, err := p.client.Pipelines.ListProjectPipelines(id.Get(), &gl.ListProjectPipelinesOptions{
		ListOptions: gl.ListOptions{},
		SHA:         gl.String(sha),
		Ref:         refP,
	})

	if err != nil {
		return pipelines, err
	}

	for _, info := range pipelinesInfos {
		if info != nil {
			pipelines = append(pipelines, *info)
		}
	}

	return pipelines, err
}

func NewPipeline(client *gl.Client) Pipeline {
	return &pipeline{client: client, project: NewProject(client)}
}
