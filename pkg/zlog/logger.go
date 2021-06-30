package zlog

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

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

func Info(format string, msgs ...interface{}) {
	log.Info().Msgf(format, msgs...)
}

func Debug(format string, msgs ...interface{}) {
	log.Debug().Msgf(format, msgs...)
}

func Warn(format string, msgs ...interface{}) {
	log.Warn().Msgf(format, msgs...)
}

func Error(format string, msgs ...interface{}) {
	log.Error().Msgf(format, msgs...)
}

func Fatal(format string, msgs ...interface{}) {
	log.Fatal().Msgf(format, msgs...)
}
