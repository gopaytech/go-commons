package asynq

import (
	"fmt"

	"github.com/gopaytech/go-commons/pkg/zlog"
)

var Logger = &LoggerAsynq{}

type LoggerAsynq struct {
}

func (a *LoggerAsynq) Debug(args ...interface{}) {
	zlog.D().Msg(fmt.Sprint(args...))
}

func (a *LoggerAsynq) Info(args ...interface{}) {
	zlog.I().Msg(fmt.Sprint(args...))
}

func (a *LoggerAsynq) Warn(args ...interface{}) {
	zlog.W().Msg(fmt.Sprint(args...))
}

func (a *LoggerAsynq) Error(args ...interface{}) {
	zlog.E().Msg(fmt.Sprint(args...))
}

func (a *LoggerAsynq) Fatal(args ...interface{}) {
	zlog.F().Msg(fmt.Sprint(args...))
}
