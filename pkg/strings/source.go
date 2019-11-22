package strings

import (
	"bytes"
	"io"
)

func FromReader(reader io.ReadCloser) (out string, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	defer reader.Close()
	out = buf.String()
	return
}
