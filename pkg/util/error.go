package util

import (
	"fmt"
	"strconv"
	"strings"
)

type CommandError error

func GetExitCode(commandError CommandError) (exitCode int) {
	if commandError != nil {
		errString := commandError.Error()
		splits := strings.Split(errString, ":")
		code := splits[0]
		intCode, err := strconv.Atoi(code)
		exitCode = intCode
		if err != nil {
			exitCode = 42
		}
		return
	}

	return
}

func ExitCodeError(exitCode int, message string) (err error) {
	return fmt.Errorf("%v: %s", exitCode, message)
}
