package docker

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Client interface {
	Build(tar io.Reader, dockerFileName string, imageName string, imageTags []string, arguments map[string]string) (messages chan *BuildResponse, err error)
	Push(imageName string, imageTags []string) (message chan *PushResponse, err chan error)
}

type dockerClient struct {
	docker     *client.Client
	registries []Registry
}

func (c *dockerClient) Build(tar io.Reader, dockerFileName string, imageName string, imageTags []string, arguments map[string]string) (messages chan *BuildResponse, err error) {
	var buildTags []string
	for _, registry := range c.registries {
		for _, tag := range imageTags {
			imageFullName := fmt.Sprintf("%s/%s:%s", registry.Endpoint, imageName, tag)
			buildTags = append(buildTags, imageFullName)
		}
	}

	authConfigs := map[string]types.AuthConfig{}
	for _, registry := range c.registries {
		authConfig := types.AuthConfig{
			Username: registry.Username,
			Password: registry.Password,
		}
		authConfigs[registry.Endpoint] = authConfig
	}

	dockerBuildOptions := types.ImageBuildOptions{
		AuthConfigs: authConfigs,
	}

	ioResult, _, err := ImageBuildWithOpts(c.docker, tar, dockerFileName, buildTags, arguments, dockerBuildOptions)
	if err != nil {
		return
	}

	messages = make(chan *BuildResponse)
	scanner := bufio.NewScanner(ioResult)
	scanner.Split(bufio.ScanLines)

	go func() {
		defer close(messages)
		for scanner.Scan() {
			msg := &BuildResponse{}
			err = json.Unmarshal(scanner.Bytes(), msg)
			if err != nil {
				msg.Error = err.Error()
			}

			messages <- msg
		}

		_ = ioResult.Close()
	}()
	return
}

func (c *dockerClient) Push(imageName string, imageTags []string) (messages chan *PushResponse, err chan error) {
	buildTags := map[string]types.AuthConfig{}

	for _, registry := range c.registries {
		for _, tag := range imageTags {
			if registry.Publish {
				imageFullName := fmt.Sprintf("%s/%s:%s", registry.Endpoint, imageName, tag)
				buildTags[imageFullName] = types.AuthConfig{
					Username: registry.Username,
					Password: registry.Password,
				}
			}
		}
	}

	messages = make(chan *PushResponse)
	err = make(chan error)
	go func() {
		defer close(messages)
		defer close(err)

		if len(buildTags) == 0 {
			err <- fmt.Errorf("no docker registries found")
			return
		}

		for tag, auth := range buildTags {
			ioResult, ierr := ImagePush(c.docker, tag, auth)
			if ierr != nil {
				err <- ierr
				return
			}

			scanner := bufio.NewScanner(ioResult)
			scanner.Split(bufio.ScanLines)

			for scanner.Scan() {
				msg := &PushResponse{}
				ierr = json.Unmarshal(scanner.Bytes(), msg)
				if ierr != nil {
					msg.Error = ierr.Error()
					err <- ierr
					break
				}

				if len(msg.Error) > 0 {
					err <- fmt.Errorf("push image error: %s", msg.Error)
					break
				}

				messages <- msg
			}

			_ = ioResult.Close()
		}
	}()

	return
}

func NewClient(docker *client.Client, registries []Registry) Client {
	return &dockerClient{docker: docker, registries: registries}
}
