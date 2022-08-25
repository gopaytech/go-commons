package gitlab

import (
	"errors"

	gl "github.com/xanzy/go-gitlab"
)

type Tag interface {
	GetLatestTag(pid NameOrId, search string) (*gl.Tag, error)
}

type tag struct {
	client *gl.Client
}

func (t *tag) GetLatestTag(pid NameOrId, search string) (*gl.Tag, error) {
	tags, _, err := t.client.Tags.ListTags(pid.ID, &gl.ListTagsOptions{
		ListOptions: gl.ListOptions{},
		Search:      &search,
	})
	if err != nil {
		return nil, err
	}

	if len(tags) == 0 {
		return nil, errors.New("no tag found")
	}

	return tags[0], nil
}

func NewTag(client *gl.Client) Tag {
	return &tag{client: client}
}
