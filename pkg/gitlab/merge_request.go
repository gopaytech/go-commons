package gitlab

import (
	gl "github.com/xanzy/go-gitlab"
)

type MergeRequestState string

func (s MergeRequestState) Equal(state string) bool {
	return string(s) == state
}

const (
	StateOpened MergeRequestState = "opened"
	StateClosed MergeRequestState = "closed"
	StateLocked MergeRequestState = "locked"
	StateMerged MergeRequestState = "merged"
)

type MergeRequest interface {
	Get(projectId NameOrId, mergeRequestId int) (*gl.MergeRequest, error)
	Create(projectId NameOrId, sourceBranch string, targetBranch string, title string) (*gl.MergeRequest, error)
	CreateToDefault(projectId NameOrId, sourceBranch string, title string) (*gl.MergeRequest, error)
	Approve(projectId NameOrId, mergeRequestID int) error
	Close(projectId NameOrId, mergeRequestID int) error
	Accept(projectId NameOrId, mergeRequestID int, removeBranch bool) (*gl.MergeRequest, error)
}

type mergeRequest struct {
	client  *gl.Client
	project Project
}

func NewMergeRequest(client *gl.Client) MergeRequest {
	return &mergeRequest{
		client:  client,
		project: NewProject(client),
	}
}

func (e *mergeRequest) Get(projectId NameOrId, mergeRequestId int) (*gl.MergeRequest, error) {
	mergeRequest, _, err := e.client.MergeRequests.GetMergeRequest(projectId.Get(), mergeRequestId, &gl.GetMergeRequestsOptions{})
	if err != nil {
		return nil, err
	}

	return mergeRequest, err
}

func (e *mergeRequest) CreateToDefault(projectId NameOrId, sourceBranch string, title string) (*gl.MergeRequest, error) {
	branch, err := e.project.GetDefaultBranch(projectId)
	if err != nil {
		return nil, err
	}

	return e.Create(projectId, sourceBranch, branch.Name, title)
}

func (e *mergeRequest) Create(projectId NameOrId, sourceBranch string, targetBranch string, title string) (*gl.MergeRequest, error) {
	mergeRequest, _, err := e.client.MergeRequests.CreateMergeRequest(projectId.Get(), &gl.CreateMergeRequestOptions{
		SourceBranch: gl.String(sourceBranch),
		TargetBranch: gl.String(targetBranch),
		Title:        &title,
	})

	if err != nil {
		return nil, err
	}

	return mergeRequest, err
}

func (e *mergeRequest) Approve(projectId NameOrId, mergeRequestID int) error {
	_, _, err := e.client.MergeRequestApprovals.ApproveMergeRequest(projectId.Get(), mergeRequestID, &gl.ApproveMergeRequestOptions{})
	return err
}

func (e *mergeRequest) Close(projectId NameOrId, mergeRequestID int) error {
	newState := StateClosed
	_, _, err := e.client.MergeRequests.UpdateMergeRequest(projectId.Get(), mergeRequestID, &gl.UpdateMergeRequestOptions{
		StateEvent: gl.String(string(newState)),
	})
	return err
}

func (e *mergeRequest) Accept(projectId NameOrId, mergeRequestID int, removeBranch bool) (*gl.MergeRequest, error) {
	mr, _, err := e.client.MergeRequests.AcceptMergeRequest(projectId.Get(), mergeRequestID, &gl.AcceptMergeRequestOptions{
		ShouldRemoveSourceBranch: gl.Bool(removeBranch),
	})
	if err != nil {
		return nil, err
	}

	return mr, nil
}
