package gitlab

import (
	gl "github.com/xanzy/go-gitlab"
)

type Job interface {
	GetByPipelineID(pid NameOrId, pipelineID int) ([]*gl.Job, error)
}

type job struct {
	client *gl.Client
}

func (j *job) GetByPipelineID(pid NameOrId, pipelineID int) ([]*gl.Job, error) {
	jobs, _, err := j.client.Jobs.ListPipelineJobs(pid.ID, pipelineID, nil)
	if err != nil {
		return nil, err
	}

	return jobs, nil
}

func NewJob(client *gl.Client) Job {
	return &job{client: client}
}
