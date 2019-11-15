package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gopaytech/go-commons/pkg/strings"
	"io"
)

func ImageBuild(client *client.Client,
	tar io.Reader,
	dockerfileName string,
	tags []string,
	args map[string]string,
) (output io.ReadCloser, osType string, err error) {
	return ImageBuildWithOpts(client, tar, dockerfileName, tags, args, types.ImageBuildOptions{})
}

func ImageBuildWithOpts(client *client.Client,
	tar io.Reader,
	dockerfileName string,
	tags []string,
	args map[string]string,
	ops types.ImageBuildOptions,
) (output io.ReadCloser, osType string, err error) {

	if !strings.IsStringEmpty(dockerfileName) {
		ops.Dockerfile = dockerfileName
	}

	if len(tags) > 0 {
		ops.Tags = tags
	}

	if len(args) > 0 {
		buildArgs := map[string]*string{}
		for key, value := range args {
			buildArgs[key] = &value
		}

		ops.BuildArgs = buildArgs
	}

	response, err := client.ImageBuild(context.Background(), tar, ops)
	if err != nil {
		return
	}

	output = response.Body
	osType = response.OSType

	return
}
