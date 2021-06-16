package exec

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestExecuteAndWaitSuccess(t *testing.T) {
	output, err := Command.ExecuteAndWait("ls", "-alh", "/")
	assert.Nil(t, err)
	assert.NotNil(t, output)
}

func TestExecuteAndWaitFailed(t *testing.T) {
	output, err := Command.ExecuteAndWait("lsx", "help")
	assert.NotNil(t, err)
	assert.Empty(t, output)
}

func TestExecute(t *testing.T) {
	stdOut, stdErr, err := Command.Execute("/usr/bin/bash", "cmd_loop_test.sh")
	assert.Nil(t, err)
	assert.NotNil(t, stdOut)
	assert.NotNil(t, stdErr)
	ScanAndCloseTimeout(stdOut, time.Second*3, func(s string) {
		t.Logf(s)
	})
}

func TestExecuteSuccess(t *testing.T) {
	stdOut, stdErr, err := Command.Execute("/usr/bin/bash", "cmd_loop_test.sh")
	assert.Nil(t, err)
	assert.NotNil(t, stdOut)
	assert.NotNil(t, stdErr)
	ScanAndClose(stdOut, func(s string) {
		t.Logf(s)
	})
}
