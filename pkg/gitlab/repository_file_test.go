package gitlab_test

import (
	"testing"

	"github.com/gopaytech/go-commons/pkg/encoding"
	"github.com/gopaytech/go-commons/pkg/gitlab"
	"github.com/stretchr/testify/assert"
)

func Test_repositoryFile_GetFileByPath(t *testing.T) {
	t.Run("Get file from branch testing should success", func(t *testing.T) {
		client, err := gitlab.NewClient("https://gitlab.com", "")
		assert.NoError(t, err)

		projectClient := gitlab.NewRepositoryFile(client)

		file, err := projectClient.GetFileByPath(gitlab.NameOrId{ID: 12967633}, "i<3Gitlab", "testing")
		assert.NoError(t, err)
		assert.NotNil(t, file)

		content, err := encoding.Base64Decode(file.Content)
		assert.NoError(t, err)
		assert.Equal(t, "God is a woman.", content)
	})
}

func Test_repositoryFile_GetRawFileByPath(t *testing.T) {
	t.Run("Get raw file from branch testing should success", func(t *testing.T) {
		client, err := gitlab.NewClient("https://gitlab.com", "")
		assert.NoError(t, err)

		projectClient := gitlab.NewRepositoryFile(client)

		fileByte, err := projectClient.GetRawFileByPath(gitlab.NameOrId{ID: 12967633}, "i<3Gitlab", "testing")
		assert.NoError(t, err)
		assert.NotNil(t, fileByte)
		assert.Equal(t, []byte("God is a woman."), fileByte)
	})
}
