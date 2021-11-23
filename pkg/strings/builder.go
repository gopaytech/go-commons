package strings

import (
	"fmt"
	"strings"
)

type Builder interface {
	Write(value string)
	Writeln(value string)
	WriteBytes(bytes []byte)
	WriteF(format string, args ...interface{})
	WriteFln(format string, args ...interface{})
	ToString() (out string)
	ToStringReset() (out string)
	Reset()
}

type builder struct {
	stringBuilder *strings.Builder
}

func (builder builder) WriteBytes(bytes []byte) {
	builder.stringBuilder.Write(bytes)
}

func (builder builder) Write(value string) {
	builder.stringBuilder.WriteString(value)
}

func (builder builder) Writeln(value string) {
	builder.stringBuilder.WriteString(value)
	builder.stringBuilder.WriteString("\n")
}

func (builder builder) WriteF(format string, args ...interface{}) {
	builder.stringBuilder.WriteString(fmt.Sprintf(format, args...))
}

func (builder builder) WriteFln(format string, args ...interface{}) {
	builder.stringBuilder.WriteString(fmt.Sprintf(format, args...))
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

func (builder builder) Reset() {
	builder.stringBuilder.Reset()
}

func NewBuilder() Builder {
	return &builder{stringBuilder: &strings.Builder{}}
}
