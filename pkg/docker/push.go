package docker

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gopaytech/go-commons/pkg/encoding"
	"io"
)

func ImagePush(client *client.Client, imageTag string, authConfig types.AuthConfig) (io.ReadCloser, error) {
	registryAuthJson, _ := json.Marshal(authConfig)
	registryAuth := encoding.Base64EncodeBytes(registryAuthJson)

	imagePushOption := types.ImagePushOptions{
		RegistryAuth: registryAuth,
	}

	return client.ImagePush(context.Background(), imageTag, imagePushOption)
}
