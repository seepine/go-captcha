package common

import (
	"go-captcha/utils"
	"os"
	"time"

	"github.com/rs/zerolog"
	logger "github.com/rs/zerolog/log"
)

func InitLog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger.Logger = logger.With().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime})
	logLevel := utils.Getenv("LOG_LEVEL", "debug")
	var l = zerolog.InfoLevel
	switch logLevel {
	case "debug":
		l = zerolog.DebugLevel
	case "warn":
		l = zerolog.WarnLevel
	case "error":
		l = zerolog.ErrorLevel
	}
	zerolog.SetGlobalLevel(l)
	logger.Info().Msg("Init log [" + logLevel + "] success")
}

func Info(format string, v ...any) {
	logger.Info().Msgf(format, v...)
}

func Warn(format string, v ...any) {
	logger.Warn().Msgf(format, v...)
}

func Error(format string, v ...any) {
	logger.Error().Msgf(format, v...)
}
