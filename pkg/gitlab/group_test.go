//go:build local && integration
// +build local,integration

package gitlab_test

import (
	"github.com/gopaytech/go-commons/pkg/gitlab"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroup_GetGroup(t *testing.T) {
	t.Run("Get repository should success", func(t *testing.T) {
		groupID := 5474201
		client, err := gitlab.NewClient("https://gitlab.com", "")
		assert.NoError(t, err)
		groupClient := gitlab.NewGroup(client)

		group, err := groupClient.GetGroup(gitlab.NameOrId{ID: groupID})
		assert.NoError(t, err)
		assert.Equal(t, group.Name, "whatabit")
		assert.Equal(t, group.WebURL, "https://gitlab.com/groups/whatabit")
	})
}
