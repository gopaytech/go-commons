//go:build local && integration
// +build local,integration

package docker

import (
	"os"
	"testing"

	"github.com/docker/docker/client"
	"github.com/stretchr/testify/assert"
)

type dockerIntegrationCtx struct {
	client       Client
	dockerClient *client.Client
}

func (ctx *dockerIntegrationCtx) SetUp(t *testing.T) {
	base64RegistriesJson, available := os.LookupEnv("TEST_DOCKER_REGISTRIES")
	assert.True(t, available)
	assert.NotNil(t, base64RegistriesJson)

	dockerHost, available := os.LookupEnv("TEST_DOCKER_HOST")
	assert.True(t, available)
	assert.NotNil(t, dockerHost)

	clientWithHost, err := NewDockerWithHost(dockerHost)
	assert.Nil(t, err)
	assert.NotNil(t, clientWithHost)

	ctx.dockerClient = clientWithHost

	registries, err := ParseRegistries(base64RegistriesJson)
	assert.Nil(t, err)
	assert.NotNil(t, registries)
	ctx.client = &dockerClient{
		docker:     ctx.dockerClient,
		registries: registries,
	}
}

func (ctx *dockerIntegrationCtx) TearDown(t *testing.T) {
	err := ctx.dockerClient.Close()
	assert.Nil(t, err)
}

func TestImageBuildAndPush(t *testing.T) {
	context := dockerIntegrationCtx{}
	context.SetUp(t)
	defer context.TearDown(t)

	imageName := "test-cx-packager"
	tags := []string{"latest", "1.0.0"}
	arguments := map[string]string{
		"FILENAME": "docker_content.txt",
	}

	tarSource, err := os.Open("./docker_integration_test.tar")
	assert.Nil(t, err)

	t.Run("build docker image and push image should success", func(t *testing.T) {
		messages, err := context.client.Build(tarSource, "Dockerfile", imageName, tags, arguments)
		assert.Nil(t, err)

		for el := range messages {
			assert.NotNil(t, el)
		}

		pushMessage, errChan := context.client.Push(imageName, tags)

		for pushMessage != nil || errChan != nil {
			select {
			case msg, ok := <-pushMessage:
				if !ok {
					pushMessage = nil
				} else {
					t.Log(msg)
					assert.NotNil(t, msg)
				}
			case e, ok := <-errChan:
				if !ok {
					errChan = nil
				} else {
					t.Log(e.Error())
					assert.Nil(t, e, "this line should not ever be executed, unless push is failed")
				}
			}
		}

		assert.True(t, true)
	})
}
