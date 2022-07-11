package gitlab

import gl "github.com/xanzy/go-gitlab"

type ProtectedBranch interface {
    CreateProtectedBranch(projectID NameOrId, branchName string) (*gl.ProtectedBranch, *gl.Response, error)
}

type protectedBranch struct {
    client *gl.Client
}

func (p *protectedBranch) CreateProtectedBranch(projectID NameOrId, branchName string) (*gl.ProtectedBranch, *gl.Response, error) {
    opts := &gl.ProtectRepositoryBranchesOptions{
        Name: &branchName,
    }
    return p.client.ProtectedBranches.ProtectRepositoryBranches(projectID.ID, opts)
}

func NewProtectedBranch(client *gl.Client) ProtectedBranch {
    return &protectedBranch{client: client}
}
