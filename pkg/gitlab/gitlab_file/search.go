package gitlab_file

import (
	"bufio"
	"bytes"
	"strings"
)

type SearchFunc func(content []byte) bool

func ContainsAll(lines ...string) SearchFunc {
	return func(content []byte) bool {
		scanner := bufio.NewScanner(bytes.NewReader(content))
		scanner.Split(bufio.ScanLines)

		scanMap := map[string]bool{}
		for _, line := range lines {
			scanMap[line] = false
		}

		for scanner.Scan() {
			for _, line := range lines {
				if strings.Contains(scanner.Text(), line) {
					scanMap[line] = true
				}
			}
		}

		for _, value := range scanMap {
			if !value {
				return false
			}
		}

		return true
	}
}

func ContainsAny(lines ...string) SearchFunc {
	return func(content []byte) bool {
		scanner := bufio.NewScanner(bytes.NewReader(content))
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			for _, line := range lines {
				if strings.Contains(scanner.Text(), line) {
					return true
				}
			}
		}

		return false
	}
}
