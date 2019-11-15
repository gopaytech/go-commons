package encoding

import (
	"encoding/base32"
)

func Base32Decode(encoded string) (plain string, err error) {
	byteString, err := base32.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(byteString), nil
}

func Base32Encode(plain string) (encoded string) {
	return base32.StdEncoding.EncodeToString([]byte(plain))
}

func Base32EncodeBytes(bytes []byte) (encoded string) {
	return base32.StdEncoding.EncodeToString(bytes)
}
