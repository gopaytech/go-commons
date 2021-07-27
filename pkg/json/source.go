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

//Convert this function will marshall source to JSON and Unmarshal it back to destination
func Convert(source interface{}, destination interface{}) error {
	marshal, err := json.Marshal(source)
	if err != nil {
		return err
	}

	err = json.Unmarshal(marshal, destination)
	if err != nil {
		return err
	}
	return nil
}
