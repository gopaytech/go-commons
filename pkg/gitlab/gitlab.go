package gitlab

import (
	"strings"

	gl "github.com/xanzy/go-gitlab"
)

func NewClient(url string, token string) (*gl.Client, error) {
	return gl.NewClient(token, gl.WithBaseURL(url))
}

type NameOrId struct {
	Name string
	id   int
}

func NewName(name string) NameOrId {
	return NameOrId{Name: name}
}

func NewNameWithClient(name string, client *gl.Client) NameOrId {
	baseUrl := client.BaseURL().String()
	name = strings.TrimPrefix(name, baseUrl)
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimSuffix(name, "/")

	return NameOrId{Name: name}
}

func NewNameWithBaseUrl(name string, baseUrl string) NameOrId {
	name = strings.TrimPrefix(name, baseUrl)
	name = strings.TrimPrefix(name, "/")
	name = strings.TrimSuffix(name, "/")

	return NameOrId{Name: name}
}

func NewId(id int) NameOrId {
	return NameOrId{id: id}
}

func (ni *NameOrId) Get() interface{} {
	if len(ni.Name) > 0 {
		return ni.Name
	}

	return ni.id
}
