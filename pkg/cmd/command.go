package cmd

import (
	"bufio"
	"fmt"
	"github.com/gopaytech/go-commons/pkg/file"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/gopaytech/go-commons/pkg/strings"
)

type CommandInterface interface {
	Exec(command string, arg ...string) (cmd *exec.Cmd, stdOut io.ReadCloser, stdErr io.ReadCloser, err error)
	ExecAndWait(command string, arg ...string) (combinedOutput string, err error)

	Execute(env map[string]string, dir string, command string, arg ...string) (cmd *exec.Cmd, stdOut io.ReadCloser, stdErr io.ReadCloser, err error)
	ExecuteAndWait(env map[string]string, dir string, command string, arg ...string) (combinedOutput string, err error)
}

type command struct{}

//ExecAndWait Simple execute and Wait
func (c *command) ExecAndWait(command string, arg ...string) (combinedOutput string, err error) {
	return c.ExecuteAndWait(map[string]string{}, "", command, arg...)
}

//ExecuteAndWait Execute command with env and working directory and wait until command exit
func (c *command) ExecuteAndWait(env map[string]string, dir string, command string, arg ...string) (combinedOutput string, err error) {
	cmd := exec.Command(command, arg...)
	envs := strings.MapKVJoin(env)
	cmd.Env = append(os.Environ(), envs...)

	if !strings.IsStringEmpty(dir) {
		if !file.DirExists(dir) {
			return "", fmt.Errorf("working directory %s is not exists or readable", dir)
		}
		cmd.Dir = dir
	}

	byteOutput, err := cmd.CombinedOutput()

	if len(byteOutput) > 0 {
		combinedOutput = string(byteOutput)
	}

	if err != nil {
		return
	}

	return
}

//Execute command with env and working directory, this func will not block
func (c *command) Execute(env map[string]string, dir string, command string, arg ...string) (cmd *exec.Cmd, stdOut io.ReadCloser, stdErr io.ReadCloser, err error) {
	cmd = exec.Command(command, arg...)
	envs := strings.MapKVJoin(env)
	cmd.Env = append(os.Environ(), envs...)

	if !strings.IsStringEmpty(dir) {
		if !file.DirExists(dir) {
			return nil, nil, nil, fmt.Errorf("working directory %s is not exists or readable", dir)
		}
		cmd.Dir = dir
	}

	stdOut, _ = cmd.StdoutPipe()
	stdErr, _ = cmd.StderrPipe()

	err = cmd.Start()

	return
}

//Exec Simple execute, this func will not block
func (c *command) Exec(command string, arg ...string) (cmd *exec.Cmd, stdOut io.ReadCloser, stdErr io.ReadCloser, err error) {
	return c.Execute(map[string]string{}, "", command, arg...)
}

//PipeToFile create or truncate file
func PipeToFile(out io.ReadCloser, path string) (err error) {
	defer out.Close()
	openFile, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return
	}
	defer openFile.Close()
	_, err = openFile.ReadFrom(out)
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
