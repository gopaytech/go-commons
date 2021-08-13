package docker

import (
	"encoding/json"

	"github.com/gopaytech/go-commons/pkg/encoding"
)

type Registry struct {
	Endpoint string `json:"endpoint"`
	Username string `json:"username"`
	Password string `json:"password"`
	Publish  bool   `json:"publish"`
}

func ParseRegistries(base64Json string) (registries []Registry, err error) {
	jsonPlain, err := encoding.Base64Decode(base64Json)
	if err != nil {
		return
	}
	registries = []Registry{}
	err = json.Unmarshal([]byte(jsonPlain), &registries)
	return
}

type BuildError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BuildResponse struct {
	Stream      string      `json:"stream"`
	ErrorDetail *BuildError `json:"errorDetail"`
	Error       string      `json:"error"`
}

type PushError struct {
	Message string `json:"message"`
}

type PushResponse struct {
	ErrorDetail *PushError `json:"errorDetail"`
	Error       string     `json:"error"`
	Status      string     `json:"status"`
	Id          string     `json:"id"`
}
