package zlog

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type LogFields map[string]interface{}

func Initialize(debug bool) {
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger().Level(zerolog.GlobalLevel())
}

func SetPrettyOutput() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log.Logger = zerolog.New(output).With().Timestamp().Logger().Level(zerolog.GlobalLevel())
}

func I() *zerolog.Event {
	return log.Info()
}

func D() *zerolog.Event {
	return log.Debug()
}

func W() *zerolog.Event {
	return log.Warn()
}

func E() *zerolog.Event {
	return log.Error()
}

func F() *zerolog.Event {
	return log.Fatal()
}

func Info(format string, msgs ...interface{}) {
	log.Info().Msgf(format, msgs...)
}

func Debug(format string, msgs ...interface{}) {
	log.Debug().Msgf(format, msgs...)
}

func Warn(format string, msgs ...interface{}) {
	log.Warn().Msgf(format, msgs...)
}

func Error(err error, format string, msgs ...interface{}) {
	log.Error().Stack().Err(err).Msgf(format, msgs...)
}

func Fatal(format string, msgs ...interface{}) {
	log.Fatal().Msgf(format, msgs...)
}

func InfoF(fields LogFields, format string, msgs ...interface{}) {
	log.Info().Fields(fields).Msgf(format, msgs...)
}

func DebugF(fields LogFields, format string, msgs ...interface{}) {
	log.Debug().Fields(fields).Msgf(format, msgs...)
}

func WarnF(fields LogFields, format string, msgs ...interface{}) {
	log.Warn().Fields(fields).Msgf(format, msgs...)
}

func ErrorF(fields LogFields, err error, format string, msgs ...interface{}) {
	log.Error().Fields(fields).Stack().Err(err).Msgf(format, msgs...)
}

func FatalF(fields LogFields, format string, msgs ...interface{}) {
	log.Fatal().Fields(fields).Msgf(format, msgs...)
}

func Ctx(ctx context.Context) *zerolog.Logger {
	return log.Ctx(ctx)
}
