package cmd

import (
	"log"
	"os"
	"os/exec"
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
