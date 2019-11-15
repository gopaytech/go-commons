package docker

import (
	"fmt"
	"github.com/docker/docker/client"
	"github.com/gopaytech/go-commons/pkg/strings"
	"io/ioutil"
	"log"
	"net/http"
	"time"
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

func NewDefaultClient() (dockerClient *client.Client, err error) {
	return NewClientWithConfig(ClientConfig{})
}

func NewClient(host string) (dockerClient *client.Client, err error) {
	return NewClientWithConfig(ClientConfig{
		Host: host,
	})
}

func NewClientWithConfig(clientConfig ClientConfig) (dockerClient *client.Client, err error) {
	dockerClient, err = client.NewClientWithOpts(clientConfig.DockerOps()...)
	return
}
