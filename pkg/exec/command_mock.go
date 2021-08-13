package exec

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type CommandMock struct {
	mock.Mock
}

func (c *CommandMock) ExecuteAndWait(command string, arg ...string) (combinedOutput string, err error) {
	arguments := c.Called(command, arg)
	return arguments.Get(0).(string), arguments.Error(1)
}

func (c *CommandMock) Execute(command string, arg ...string) (stdOut io.ReadCloser, stdErr io.ReadCloser, err error) {
	arguments := c.Called(command, arg)
	return arguments.Get(0).(io.ReadCloser), arguments.Get(1).(io.ReadCloser), arguments.Error(2)
}

func (c *CommandMock) ExecuteAndWaitWithEnv(env map[string]string, command string, arg ...string) (combinedOutput string, err error) {
	arguments := c.Called(env, command, arg)
	return arguments.Get(0).(string), arguments.Error(1)
}

func (c *CommandMock) ExecuteWithEnv(env map[string]string, command string, arg ...string) (stdOut io.ReadCloser, stdErr io.ReadCloser, err error) {
	arguments := c.Called(env, command, arg)
	return arguments.Get(0).(io.ReadCloser), arguments.Get(1).(io.ReadCloser), arguments.Error(2)
}
