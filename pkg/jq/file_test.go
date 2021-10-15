package jq

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileQuery_Execute(t *testing.T) {
	t.Run("given object json should success", func(t *testing.T) {
		query := &FileQuery{
			Path:      "./test/posts.json",
			Query:     ".data[].title",
			QueryType: ObjectType,
		}

		err := query.Execute(context.Background(), func(value interface{}) bool {
			assert.NotEmpty(t, value)
			return true
		})

		assert.NoError(t, err)
	})

	t.Run("given array json should success", func(t *testing.T) {
		arrayQuery := NewFileQuery("./test/post-array.json", ".[].body", ListType)
		err := arrayQuery.Execute(context.Background(), func(value interface{}) bool {
			assert.NotEmpty(t, value)
			return true
		})

		assert.NoError(t, err)
	})

	t.Run("given non existing path should failed", func(t *testing.T) {
		query := &FileQuery{
			Path:      "./test/nonexistence.json",
			Query:     ".data[].title",
			QueryType: ObjectType,
		}

		err := query.Execute(context.Background(), func(value interface{}) bool {
			// should not be called
			assert.Empty(t, value)
			return false
		})

		assert.Error(t, err)
	})
}
