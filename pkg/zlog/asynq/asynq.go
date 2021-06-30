package asynq

import (
	"github.com/gopaytech/go-commons/pkg/zlog"
	"github.com/spf13/cast"
)

type Logger struct {
}

func splitArgs(args ...interface{}) (format string, arguments []interface{}) {
	if len(args) <= 0 {
		return "", []interface{}{}
	}
	if len(args) == 1 {
		return "%s", []interface{}{cast.ToString(args[0])}
	}

	return cast.ToString(args[0]), args[1:]
}

func (a *Logger) Debug(args ...interface{}) {
	zlog.Debug(splitArgs(args))
}

func (a *Logger) Info(args ...interface{}) {
	zlog.Info(splitArgs(args))
}

func (a *Logger) Warn(args ...interface{}) {
	zlog.Warn(splitArgs(args))
}

func (a *Logger) Error(args ...interface{}) {
	zlog.Error(splitArgs(args))
}

func (a *Logger) Fatal(args ...interface{}) {
	zlog.Fatal(splitArgs(args))
}
