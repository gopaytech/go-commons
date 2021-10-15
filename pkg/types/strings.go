package types

import (
	"fmt"
	"strings"
)

func ToKeyValueMap(origin string) map[string]string {
	entries := strings.Split(origin, ";")

	result := make(map[string]string)
	for _, e := range entries {
		parts := strings.Split(e, "=")
		if len(parts) >= 2 {
			result[parts[0]] = parts[1]
		}
	}

	return result
}

func FromKeyValueMap(mapOrigin map[string]string) (string, error) {
	if len(mapOrigin) == 0 {
		return "", nil
	}

	var builder strings.Builder
	for key, value := range mapOrigin {
		_, err := fmt.Fprintf(&builder, "%s=%s;", key, value)
		if err != nil {
			return "", err
		}
	}
	result := builder.String()
	if builder.Len() > 1 {
		result = result[:builder.Len()-1]
	}

	return result, nil

}
