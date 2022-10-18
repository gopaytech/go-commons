package gitlab_test

import (
	"testing"

	"github.com/gopaytech/go-commons/pkg/gitlab"
	"github.com/stretchr/testify/assert"
)

func Test_tag_GetLatestTag(t *testing.T) {
	t.Run("Get file from branch testing should success", func(t *testing.T) {
		client, err := gitlab.NewClient("https://gitlab.com", "")
		assert.NoError(t, err)

		projectClient := gitlab.NewTag(client)

		tag, err := projectClient.GetLatestTag(gitlab.NameOrId{ID: 17954603}, "^v")
		assert.NoError(t, err)
		assert.NotNil(t, tag)
		assert.Equal(t, "v1.0.0", tag.Name)
	})
}
