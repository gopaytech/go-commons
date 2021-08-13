package gorm

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

var (
	infoStr     = "%s: [info] "
	warnStr     = "%s: [warn] "
	errStr      = "%s: [error] "
	traceStr    = "%s: [%.3fms] [rows:%v] %s"
	traceErrStr = "%s %s [%.3fms] [rows:%v] %s"
)

var GormLogger = Logger{}

type Logger struct {
}

func (l Logger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l Logger) Error(ctx context.Context, msg string, opts ...interface{}) {
	zctx := log.Logger.WithContext(ctx)
	log.Ctx(zctx).Error().Msgf(errStr+msg, append([]interface{}{utils.FileWithLineNum()}, opts...)...)
}

func (l Logger) Warn(ctx context.Context, msg string, opts ...interface{}) {
	zctx := log.Logger.WithContext(ctx)
	log.Ctx(zctx).Warn().Msgf(warnStr+msg, append([]interface{}{utils.FileWithLineNum()}, opts...)...)
}

func (l Logger) Info(ctx context.Context, msg string, opts ...interface{}) {
	zctx := log.Logger.WithContext(ctx)
	log.Ctx(zctx).Info().Msgf(infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, opts...)...)
}

func (l Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	zctx := log.Logger.WithContext(ctx)
	zlog := log.Ctx(zctx)

	logLevel := zlog.GetLevel()

	elapsed := time.Since(begin)

	if logLevel == zerolog.DebugLevel {
		if err != nil {
			sql, rows := fc()
			if rows == -1 {
				zlog.Printf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				zlog.Printf(traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		} else {
			sql, rows := fc()
			if rows == -1 {
				zlog.Printf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				zlog.Printf(traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}
