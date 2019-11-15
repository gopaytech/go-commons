package dir

import (
	"fmt"
	"os"
	"os/user"
)

func Exists(filePath string) bool {
	info, err := os.Stat(filePath)
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
