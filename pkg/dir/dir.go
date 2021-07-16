package dir

import (
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func IsEmpty(name string) bool {
	d, err := os.Stat(name)
	if err != nil {
		return false
	}

	if !d.IsDir() {
		return false
	}

	f, err := os.Open(name)
	if err != nil {
		return false
	}

	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true
	}

	return false
}

func Exists(dirPath string) bool {
	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func Home(children ...string) string {
	usr, _ := user.Current()

	home := usr.HomeDir
	for _, child := range children {
		home = fmt.Sprintf("%s/%s", home, child)
	}

	return home
}

func ListFiles(directory string) (files []string, err error) {
	baseDirectory := filepath.Clean(directory)
	err = filepath.Walk(baseDirectory,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			cleanPath := strings.Replace(path, fmt.Sprintf("%s/", baseDirectory), "", 1)
			files = append(files, cleanPath)
			return nil
		})
	return
}
