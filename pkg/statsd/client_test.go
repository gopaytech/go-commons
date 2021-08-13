package statsd_test

import (
	"testing"

	"github.com/gopaytech/go-commons/pkg/statsd"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	config := statsd.ClientConfig{
		Host:     "localhost:8000",
		Protocol: "tcp",
		Prefix:   "packager",
		Muted:    true,
	}
	statsdClient, err := statsd.NewClient(config)

	assert.NoError(t, err)
	assert.NotNil(t, statsdClient)
}
