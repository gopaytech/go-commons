package strings

import (
	"fmt"
	"strings"
)

type Builder interface {
	Write(format string, args ...interface{})
	Writeln(format string, args ...interface{})
	ToString() (out string)
	ToStringReset() (out string)
}

type builder struct {
	stringBuilder strings.Builder
}

func (builder builder) Write(format string, args ...interface{}) {
	builder.stringBuilder.WriteString(fmt.Sprintf(format, args))
}

func (builder builder) Writeln(format string, args ...interface{}) {
	builder.stringBuilder.WriteString(fmt.Sprintf(format, args))
	builder.stringBuilder.WriteString("\n")
}

func (builder builder) ToString() (out string) {
	return builder.stringBuilder.String()
}

func (builder builder) ToStringReset() (out string) {
	out = builder.stringBuilder.String()
	builder.stringBuilder.Reset()
	return
}

func NewBuilder() Builder {
	var stringBuilder strings.Builder
	return &builder{stringBuilder: stringBuilder}
}
