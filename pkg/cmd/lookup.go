package cmd

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

func LookPathDefault(executableName string, defaultPath string) (binaryPath string) {
	path, err := exec.LookPath(executableName)
	if err != nil {
		return defaultPath
	}
	return path
}

func LookPath(executableName string) (binaryPath string) {
	path, err := exec.LookPath(executableName)
	if err != nil {
		log.Panicln(err.Error())
	}
	return path
}

func LookupEnvDefault(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func LookupEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Panicf("env key %s not found\n", key)
	}
	return value
}

func OsEnv() map[string]string {
	output := map[string]string{}
	for _, env := range os.Environ() {
		envPair := strings.SplitN(env, "=", 2)
		key := envPair[0]
		value := envPair[1]
		output[key] = value
	}
	return output
}
