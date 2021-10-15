package jq

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteQuery_Execute(t *testing.T) {
	server := httptest.NewServer(
		http.FileServer(http.Dir("test")),
	)

	defer server.Close()

	t.Run("given object json should pass", func(t *testing.T) {
		query := &RemoteQuery{
			SourceUrl: fmt.Sprintf("%s/posts.json", server.URL),
			Query:     ".data[].title",
			QueryType: ObjectType,
		}

		err := query.Execute(context.Background(), func(value interface{}) bool {
			assert.NotEmpty(t, value)
			return true
		})

		assert.NoError(t, err)
	})

	t.Run("given array json should pass", func(t *testing.T) {
		arrayQuery := NewRemoteQuery(fmt.Sprintf("%s/post-array.json", server.URL), map[string]string{}, ".[].body", ListType)
		err := arrayQuery.Execute(context.Background(), func(value interface{}) bool {
			assert.NotEmpty(t, value)
			return true
		})

		assert.NoError(t, err)
	})

	t.Run("given non exists url should fail", func(t *testing.T) {
		query := &RemoteQuery{
			SourceUrl: "httpx://randomsite.org",
			Headers: map[string]string{
				"title": "something",
			},
			Query:     ".data[].title",
			QueryType: ObjectType,
		}

		err := query.Execute(context.Background(), func(value interface{}) bool {
			assert.NotEmpty(t, value)
			return true
		})

		assert.Error(t, err)
	})
}
