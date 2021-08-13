package exec

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/fatih/color"
	"github.com/gopaytech/go-commons/pkg/stdout"
	"github.com/gopaytech/go-commons/pkg/strings"
)

type CommandInterface interface {
	ExecuteAndWait(command string, arg ...string) (combinedOutput string, err error)
	Execute(command string, arg ...string) (stdOut io.ReadCloser, stdErr io.ReadCloser, err error)
	ExecuteWithEnv(env map[string]string, command string, arg ...string) (stdOut io.ReadCloser, stdErr io.ReadCloser, err error)
	ExecuteAndWaitWithEnv(env map[string]string, command string, arg ...string) (combinedOutput string, err error)
}

type command struct{}

func (c *command) ExecuteAndWait(command string, arg ...string) (combinedOutput string, err error) {
	stdout.ColorPrinter.Println(color.FgHiBlue, "%v %v", command, arg)
	cmd := exec.Command(command, arg...)
	byteOutput, err := cmd.CombinedOutput()

	if len(byteOutput) > 0 {
		combinedOutput = string(byteOutput)
	}

	if err != nil {
		return
	}

	return
}

func (c *command) ExecuteAndWaitWithEnv(env map[string]string, command string, arg ...string) (combinedOutput string, err error) {
	stdout.ColorPrinter.Println(color.FgHiBlue, "%v %v %v", strings.MapKVJoin(env), command, arg)
	cmd := exec.Command(command, arg...)
	envs := strings.MapKVJoin(env)
	cmd.Env = append(os.Environ(), envs...)

	byteOutput, err := cmd.CombinedOutput()

	if len(byteOutput) > 0 {
		combinedOutput = string(byteOutput)
	}

	if err != nil {
		return
	}

	return
}

func (c *command) ExecuteWithEnv(env map[string]string, command string, arg ...string) (stdOut io.ReadCloser, stdErr io.ReadCloser, err error) {
	stdout.ColorPrinter.Println(color.FgHiBlue, "%v %v %v", strings.MapKVJoin(env), command, arg)
	cmd := exec.Command(command, arg...)
	envs := strings.MapKVJoin(env)
	cmd.Env = append(os.Environ(), envs...)
	stdOut, _ = cmd.StdoutPipe()
	stdErr, _ = cmd.StderrPipe()

	err = cmd.Start()

	return
}

func (c *command) Execute(command string, arg ...string) (stdOut io.ReadCloser, stdErr io.ReadCloser, err error) {
	stdout.ColorPrinter.Println(color.FgHiBlue, "%v %v", command, arg)
	cmd := exec.Command(command, arg...)
	stdOut, _ = cmd.StdoutPipe()
	stdErr, _ = cmd.StderrPipe()

	err = cmd.Start()

	return
}

func ScanAndClose(out io.ReadCloser, ops func(string)) {
	defer out.Close()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		ops(string(scanner.Bytes()))
	}
}

func ScanAndCloseTimeout(out io.ReadCloser, timeout time.Duration, ops func(string)) {
	defer out.Close()

	scanner := bufio.NewScanner(out)
	scanner.Split(bufio.ScanLines)

	timeoutChan := time.After(timeout)
	valueChan := make(chan string, 1)

	go func() {
		defer close(valueChan)
		for scanner.Scan() {
			valueChan <- string(scanner.Bytes())
		}
	}()

	for {
		select {
		case value, ok := <-valueChan:
			if ok {
				ops(value)
				break
			} else {
				return
			}
		case <-timeoutChan:
			return
		}
	}
}

func NewCommand() CommandInterface {
	return &command{}
}

var Command = NewCommand()
