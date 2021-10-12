package cmd

import (
	"fmt"
	"github.com/gopaytech/go-commons/pkg/strings"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExecuteAndWaitSuccess(t *testing.T) {
	output, err := Command.ExecAndWait("ls", "-alh", "/")
	assert.Nil(t, err)
	assert.NotNil(t, output)
}

func TestExecuteAndWaitFailed(t *testing.T) {
	output, err := Command.ExecAndWait("lsx", "help")
	assert.NotNil(t, err)
	assert.Empty(t, output)
}

func TestExecute(t *testing.T) {
	_, stdOut, stdErr, err := Command.Exec("/usr/bin/bash", "cmd_loop_test.sh")
	assert.Nil(t, err)
	assert.NotNil(t, stdOut)
	assert.NotNil(t, stdErr)
	ScanAndCloseTimeout(stdOut, time.Second*3, func(s string) {
		t.Logf(s)
	})
}

func TestExecuteSuccess(t *testing.T) {
	_, stdOut, stdErr, err := Command.Exec("/usr/bin/bash", "cmd_loop_test.sh")
	assert.Nil(t, err)
	assert.NotNil(t, stdOut)
	assert.NotNil(t, stdErr)
	ScanAndClose(stdOut, func(s string) {
		t.Logf(s)
	})
}

func TestExecuteToFile(t *testing.T) {
	_, stdOut, stdErr, err := Command.Exec("/usr/bin/bash", "cmd_loop_test.sh")
	assert.Nil(t, err)
	assert.NotNil(t, stdOut)
	assert.NotNil(t, stdErr)
	path := fmt.Sprintf("%s/%s", os.TempDir(), strings.RandomAlphanumeric(13))
	err = PipeToFile(stdOut, path)
	assert.NoError(t, err)
	assert.FileExists(t, path)
}
