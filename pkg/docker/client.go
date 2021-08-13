package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/gopaytech/go-commons/pkg/config"
	"github.com/gopaytech/go-commons/pkg/encoding"
	"github.com/gopaytech/go-commons/pkg/strings"
)

// path will used by default unless data is present
type ClientTlsConfig struct {
	caCertData string
	caCertPath string

	certData string
	certPath string

	keyData string
	keyPath string
}

type ClientConfig struct {
	Host       string
	HttpClient *http.Client
	HttpHeader map[string]string
	timeout    *time.Duration
	Tls        *ClientTlsConfig
}

func (c *ClientConfig) DockerOps() (ops []client.Opt) {
	ops = []client.Opt{
		client.WithAPIVersionNegotiation(),
	}

	if !strings.IsStringEmpty(c.Host) {
		ops = append(ops, client.WithHost(c.Host))
	}

	if c.HttpClient != nil {
		ops = append(ops, client.WithHTTPClient(c.HttpClient))
	}

	if c.timeout != nil {
		ops = append(ops, client.WithTimeout(*c.timeout))
	}

	if len(c.HttpHeader) >= 0 {
		ops = append(ops, client.WithHTTPHeaders(c.HttpHeader))
	}

	if c.Tls != nil {
		tls := c.Tls

		if !strings.IsStringEmpty(tls.caCertData) {

			tempFile, err := ioutil.TempFile("/tmp", "ca-cert")
			if err != nil {
				log.Fatal(fmt.Sprintf("cannot create tmp file to store ca-cert %s", err.Error()))
			}

			tls.caCertPath = tempFile.Name()
		}

		if !strings.IsStringEmpty(tls.certData) {

			tempFile, err := ioutil.TempFile("/tmp", "cert")
			if err != nil {
				log.Fatal(fmt.Sprintf("cannot create tmp file to store cert %s", err.Error()))
			}

			tls.certPath = tempFile.Name()
		}

		if !strings.IsStringEmpty(tls.keyData) {

			tempFile, err := ioutil.TempFile("/tmp", "key")
			if err != nil {
				log.Fatal(fmt.Sprintf("cannot create tmp file to store key %s", err.Error()))
			}

			tls.keyPath = tempFile.Name()
		}

		ops = append(ops, client.WithTLSClientConfig(tls.caCertPath, tls.certPath, tls.keyPath))
	}

	return
}

func NewDefaultDocker() (dockerClient *client.Client, err error) {
	return NewDockerWithConfig(ClientConfig{})
}

func NewDockerWithHost(host string) (dockerClient *client.Client, err error) {
	return NewDockerWithConfig(ClientConfig{
		Host: host,
	})
}

func NewDockerWithConfig(clientConfig ClientConfig) (dockerClient *client.Client, err error) {
	dockerClient, err = client.NewClientWithOpts(clientConfig.DockerOps()...)
	return
}

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
			buildArgs[key] = config.String(value)
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

func ImagePush(client *client.Client, imageTag string, authConfig types.AuthConfig) (io.ReadCloser, error) {
	registryAuthJson, _ := json.Marshal(authConfig)
	registryAuth := encoding.Base64EncodeBytes(registryAuthJson)

	imagePushOption := types.ImagePushOptions{
		RegistryAuth: registryAuth,
	}

	return client.ImagePush(context.Background(), imageTag, imagePushOption)
}
