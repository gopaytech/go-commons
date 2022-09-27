package gitlab_test

import (
	"testing"

	"github.com/gopaytech/go-commons/pkg/gitlab"
	"github.com/stretchr/testify/assert"
)

func Test_project_GetDefaultBranch(t *testing.T) {
	t.Run("Get branch master should success", func(t *testing.T) {
		client, err := gitlab.NewClient("https://gitlab.com", "")
		assert.NoError(t, err)

		projectClient := gitlab.NewProject(client)

		branch, err := projectClient.GetDefaultBranch(gitlab.NameOrId{ID: 12967633})
		assert.NoError(t, err)
		assert.Equal(t, branch.Name, "master")
		assert.Equal(t, branch.WebURL, "https://gitlab.com/whatabit/testproj/-/tree/master")
	})
}

func Test_project_GetBranchByName(t *testing.T) {
	t.Run("Get branch master should success", func(t *testing.T) {
		client, err := gitlab.NewClient("https://gitlab.com", "")
		assert.NoError(t, err)

		projectClient := gitlab.NewProject(client)

		branch, err := projectClient.GetBranchByName(gitlab.NameOrId{ID: 12967633}, "master")
		assert.NoError(t, err)
		assert.Equal(t, branch.Name, "master")
		assert.Equal(t, branch.WebURL, "https://gitlab.com/whatabit/testproj/-/tree/master")
	})
}
