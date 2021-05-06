package encoding

import "encoding/base64"

func Base64Decode(encoded string) (plain string, err error) {
	byteString, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	return string(byteString), nil
}

func Base64DecodeBytes(encoded string) (plain []byte, err error) {
	byteString, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return
	}

	return byteString, nil
}

func Base64Encode(plain string) (encoded string) {
	return base64.StdEncoding.EncodeToString([]byte(plain))
}

func Base64EncodeBytes(bytes []byte) (encoded string) {
	return base64.StdEncoding.EncodeToString(bytes)
}
