package util

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func HomeDir(children ...string) string {
	usr, _ := user.Current()

	home := usr.HomeDir
	for _, child := range children {
		home = fmt.Sprintf("%s/%s", home, child)
	}

	return home
}

func ReadFile(fileLocation string) (file io.ReadCloser, err error) {
	return os.Open(fileLocation)
}

func ReadHomeFile(children ...string) (file io.ReadCloser, err error) {
	usr, _ := user.Current()
	home := usr.HomeDir
	for _, child := range children {
		home = fmt.Sprintf("%s/%s", home, child)
	}

	return os.Open(home)
}

func ReadStdIn() (stdIn io.ReadCloser) {
	return os.Stdin
}

func ReadUrl(url string) (stream io.ReadCloser, err error) {
	response, err := http.Get(url)
	if err != nil {
		return
	}

	if response.StatusCode < 200 && response.StatusCode > 299 {
		err = fmt.Errorf("status code is error: %v", response.StatusCode)
		return
	}

	stream = response.Body
	return
}

func SaveStringToFile(source string, fileName string) (err error) {
	sourceByte := []byte(source)
	err = ioutil.WriteFile(fileName, sourceByte, 0644)
	return
}

func ReadFileToString(filePath string) (result string, err error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	result = string(content)
	return
}
