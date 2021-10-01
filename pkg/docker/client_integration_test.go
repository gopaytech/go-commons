//go:build local && integration
// +build local,integration

package docker

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type clientIntegrationTestContext struct {
}

func (ctx *clientIntegrationTestContext) SetUp(t *testing.T) {
}

func (ctx *clientIntegrationTestContext) TearDown(t *testing.T) {

}

func TestDefaultClient(t *testing.T) {
	context := clientIntegrationTestContext{}
	context.SetUp(t)
	defer context.TearDown(t)

	client, err := NewDefaultDocker()
	assert.Nil(t, err)
	assert.NotNil(t, client)
	fmt.Println(client.ClientVersion())
}

func TestClient(t *testing.T) {
	host, available := os.LookupEnv("TEST_DOCKER_HOST")
	assert.True(t, available)
	assert.NotNil(t, host)

	client, err := NewDockerWithHost(host)
	assert.Nil(t, err)
	assert.NotNil(t, client)
	fmt.Println(client.ClientVersion())
}
