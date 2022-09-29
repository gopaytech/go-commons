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

	MasterBranchName = "master"
)

type MergeRequest interface {
	Get(projectID NameOrId, mergeRequestID int) (*gl.MergeRequest, error)
	Create(projectID NameOrId, sourceBranch string, targetBranch string, title string) (*gl.MergeRequest, error)
	CreateToDefault(projectID NameOrId, sourceBranch string, title string) (*gl.MergeRequest, error)
	CreateToMaster(projectID NameOrId, sourceBranch string, title string) (*gl.MergeRequest, error)
	Approve(projectID NameOrId, mergeRequestID int) error
	Close(projectID NameOrId, mergeRequestID int) error
	Accept(projectID NameOrId, mergeRequestID int, removeBranch bool, whenPipelinePassed bool) (*gl.MergeRequest, error)
	ResetApproval(projectID NameOrId, mergeRequestID int) error
}

type mergeRequest struct {
	client  *gl.Client
	project Project
}

func (e *mergeRequest) ResetApproval(projectID NameOrId, mergeRequestID int) error {
	rules, _, err := e.client.MergeRequestApprovals.GetApprovalRules(projectID.Get(), mergeRequestID)
	if err != nil {
		return err
	}

	for _, rule := range rules {
		_, err := e.client.MergeRequestApprovals.DeleteApprovalRule(projectID.Get(), mergeRequestID, rule.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewMergeRequest(client *gl.Client) MergeRequest {
	return &mergeRequest{
		client:  client,
		project: NewProject(client),
	}
}

func (e *mergeRequest) Get(projectID NameOrId, mergeRequestID int) (*gl.MergeRequest, error) {
	mergeRequest, _, err := e.client.MergeRequests.GetMergeRequest(projectID.Get(), mergeRequestID, &gl.GetMergeRequestsOptions{})
	if err != nil {
		return nil, err
	}

	return mergeRequest, err
}

func (e *mergeRequest) CreateToDefault(projectID NameOrId, sourceBranch string, title string) (*gl.MergeRequest, error) {
	branch, err := e.project.GetDefaultBranch(projectID)
	if err != nil {
		return nil, err
	}

	return e.Create(projectID, sourceBranch, branch.Name, title)
}

func (e *mergeRequest) CreateToMaster(projectID NameOrId, sourceBranch string, title string) (*gl.MergeRequest, error) {
	branch, err := e.project.GetBranchByName(projectID, MasterBranchName)
	if err != nil {
		return nil, err
	}

	return e.Create(projectID, sourceBranch, branch.Name, title)
}

func (e *mergeRequest) Create(projectID NameOrId, sourceBranch string, targetBranch string, title string) (*gl.MergeRequest, error) {
	mergeRequest, _, err := e.client.MergeRequests.CreateMergeRequest(projectID.Get(), &gl.CreateMergeRequestOptions{
		SourceBranch: gl.String(sourceBranch),
		TargetBranch: gl.String(targetBranch),
		Title:        &title,
	})
	if err != nil {
		return nil, err
	}

	return mergeRequest, err
}

func (e *mergeRequest) Approve(projectID NameOrId, mergeRequestID int) error {
	_, _, err := e.client.MergeRequestApprovals.ApproveMergeRequest(projectID.Get(), mergeRequestID, &gl.ApproveMergeRequestOptions{})
	return err
}

func (e *mergeRequest) Close(projectID NameOrId, mergeRequestID int) error {
	newState := StateClosed
	_, _, err := e.client.MergeRequests.UpdateMergeRequest(projectID.Get(), mergeRequestID, &gl.UpdateMergeRequestOptions{
		StateEvent: gl.String(string(newState)),
	})
	return err
}

func (e *mergeRequest) Accept(projectID NameOrId, mergeRequestID int, removeBranch bool, whenPipelinePassed bool) (*gl.MergeRequest, error) {
	mr, _, err := e.client.MergeRequests.AcceptMergeRequest(projectID.Get(), mergeRequestID, &gl.AcceptMergeRequestOptions{
		ShouldRemoveSourceBranch:  gl.Bool(removeBranch),
		MergeWhenPipelineSucceeds: gl.Bool(whenPipelinePassed),
	})
	if err != nil {
		return nil, err
	}

	return mr, nil
}
