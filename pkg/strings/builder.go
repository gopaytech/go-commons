package strings

import "strings"

type Builder interface {
	WriteString(format string, args ...interface{})
	ToString() (out string)
	ToStringReset() (out string)
}

type builder struct {
	stringBuilder strings.Builder
}

func (b builder) WriteString(format string, args ...interface{}) {
	panic("implement me")
}

func (b builder) ToString() (out string) {
	return b.stringBuilder.String()
}

func (b builder) ToStringReset() (out string) {
	out = b.stringBuilder.String()
	b.stringBuilder.Reset()
	return
}

func NewBuilder() Builder {
	var stringBuilder strings.Builder
	return &builder{stringBuilder: stringBuilder}
}
