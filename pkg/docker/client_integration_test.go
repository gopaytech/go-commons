package docker

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type clientIntegrationTestContext struct {
}

func (ctx *clientIntegrationTestContext) SetUp(t *testing.T) {
	value, available := os.LookupEnv("ENABLE_INTEGRATION_TEST")
	if available != true || value != "true" {
		t.SkipNow()
	}
}

func (ctx *clientIntegrationTestContext) TearDown(t *testing.T) {

}

func TestDefaultClient(t *testing.T) {
	context := clientIntegrationTestContext{}
	context.SetUp(t)
	defer context.TearDown(t)

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
