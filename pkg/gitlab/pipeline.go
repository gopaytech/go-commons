package gitlab

import (
	"github.com/gopaytech/go-commons/pkg/strings"
	gl "github.com/xanzy/go-gitlab"
)

type PipelineStatus string

const (
	PipelineCreated            PipelineStatus = "created"
	PipelineWaitingForResource PipelineStatus = "waiting_for_resource"
	PipelinePreparing          PipelineStatus = "preparing"
	PipelinePending            PipelineStatus = "pending"
	PipelineRunning            PipelineStatus = "running"
	PipelineSuccess            PipelineStatus = "success"
	PipelineFailed             PipelineStatus = "failed"
	PipelineCanceled           PipelineStatus = "canceled"
	PipelineSkipped            PipelineStatus = "skipped"
	PipelineManual             PipelineStatus = "manual"
	PipelineScheduled          PipelineStatus = "scheduled"
)

type Pipeline interface {
	GetBySHA(projectId NameOrId, sha string) ([]gl.PipelineInfo, error)
	GetBySHAAndRef(projectId NameOrId, sha string, ref string) ([]gl.PipelineInfo, error)
	GetBySHAOnDefault(projectId NameOrId, sha string) ([]gl.PipelineInfo, error)
}

type pipeline struct {
	client  *gl.Client
	project Project
}

func IsPipelineFinished(pipeline gl.PipelineInfo) bool {
	status := PipelineStatus(pipeline.Status)
	switch status {
	case PipelineSuccess, PipelineFailed, PipelineCanceled, PipelineSkipped:
		return true
	default:
		return false
	}
}

func (p *pipeline) GetBySHA(projectId NameOrId, sha string) ([]gl.PipelineInfo, error) {
	return p.GetBySHAAndRef(projectId, sha, "")
}

func (p *pipeline) GetBySHAOnDefault(projectId NameOrId, sha string) ([]gl.PipelineInfo, error) {
	branch, err := p.project.GetDefaultBranch(projectId)
	if err != nil {
		return nil, err
	}
	return p.GetBySHAAndRef(projectId, sha, branch.Name)

}

func (p *pipeline) GetBySHAAndRef(projectId NameOrId, sha string, ref string) ([]gl.PipelineInfo, error) {
	refP := new(string)

	if !strings.IsStringEmpty(ref) {
		*refP = ref
	}

	var pipelines []gl.PipelineInfo
	pipelinesInfos, _, err := p.client.Pipelines.ListProjectPipelines(projectId.Get(), &gl.ListProjectPipelinesOptions{
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
