package strings

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func IsStringEmpty(text string) bool {
	text = Trim(text)
	if len(text) == 0 {
		return true
	}
	return false
}

func Trim(text string) string {
	return strings.TrimSpace(text)
}

func CleanEmptyString(s []string) []string {
	var r []string
	for _, str := range s {
		if !IsStringEmpty(str) {
			r = append(r, str)
		}
	}
	return r
}

func ExecuteTemplateToString(template *template.Template, value interface{}) (output string, err error) {
	stringBuffer := bytes.NewBufferString("")
	err = template.Execute(stringBuffer, value)
	if err != nil {
		return
	}

	output = stringBuffer.String()
	return

}

func MapKVJoin(values map[string]string) (result []string) {
	result = []string{}
	for key, value := range values {
		result = append(result, KVJoin(key, value))
	}
	return
}

func KVJoin(key string, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}

func KVSplit(kvString string) (key string, value string) {
	splits := strings.Split(kvString, "=")
	key = splits[0]
	if len(splits) > 1 {
		value = splits[1]
	}
	return
}
