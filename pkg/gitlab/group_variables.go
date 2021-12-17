package gitlab

import gl "github.com/xanzy/go-gitlab"

type GroupVariables interface {
	GetVariable(gid interface{}, key string) (*gl.GroupVariable, error)
}

type groupVariables struct {
	client *gl.Client
}

func (gv *groupVariables) GetVariable(gid interface{}, key string) (*gl.GroupVariable, error) {
	result, _, err := gv.client.GroupVariables.GetVariable(gid, key)
	if err != nil {
		return &gl.GroupVariable{}, err
	}

	return result, nil
}

func NewGroupVariable(client *gl.Client) GroupVariables {
	return &groupVariables{client: client}
}
