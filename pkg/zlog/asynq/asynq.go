package asynq

import (
	"github.com/gopaytech/go-commons/pkg/zlog"
	"github.com/spf13/cast"
)

var Logger = &LoggerAsynq{}

type LoggerAsynq struct {
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

func (a *LoggerAsynq) Debug(args ...interface{}) {
	zlog.Debug(splitArgs(args))
}

func (a *LoggerAsynq) Info(args ...interface{}) {
	zlog.Info(splitArgs(args))
}

func (a *LoggerAsynq) Warn(args ...interface{}) {
	zlog.Warn(splitArgs(args))
}

func (a *LoggerAsynq) Error(args ...interface{}) {
	zlog.Error(splitArgs(args))
}

func (a *LoggerAsynq) Fatal(args ...interface{}) {
	zlog.Fatal(splitArgs(args))
}
