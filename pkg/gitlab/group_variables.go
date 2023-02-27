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
func (gv *groupVariables) GetVariables(groupID int, recursive bool) (map[string]string, error) {
	variables := make(map[string]string)
	group, _, err := gv.client.Groups.GetGroup(groupID, &gl.GetGroupOptions{})
	if err != nil {
		return variables, err
	}

	for group != nil {
		v, _, err := gv.client.GroupVariables.ListVariables(groupID, &gl.ListGroupVariablesOptions{})
		if err != nil {
			return variables, err
		}
		for _, variable := range v {
			variables[variable.Key] = variable.Value
		}

		if recursive == false {
			break
		}

		groupID = group.ParentID
		if groupID != 0 {
			group, _, err = gv.client.Groups.GetGroup(groupID, &gl.GetGroupOptions{})
			if err != nil {
				return variables, err
			}
		} else {
			group = nil
		}
	}

	return variables, nil
}

func NewGroupVariable(client *gl.Client) GroupVariables {
	return &groupVariables{client: client}
}
