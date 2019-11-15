package docker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultClient(t *testing.T) {
	client, err := NewDefaultClient()
	assert.Nil(t, err)
	assert.NotNil(t, client)
	fmt.Println(client.ClientVersion())
}

func TestClient(t *testing.T) {
	client, err := NewClient("unix:///var/run/docker.sock")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	fmt.Println(client.ClientVersion())
}
