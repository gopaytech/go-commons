package json

import (
	"bytes"
	"encoding/json"
	"io"
)

func FromReader(reader io.ReadCloser, target interface{}) (err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(reader)
	if err != nil {
		return
	}
	defer reader.Close()
	return json.Unmarshal(buf.Bytes(), target)
}
