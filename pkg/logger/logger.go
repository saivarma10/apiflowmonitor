package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	log = zerolog.New(os.Stderr).With().Timestamp().Logger()
	log = log.Level(zerolog.DebugLevel)
	log.Info().Msg("Logger initialized")
}

func GetLogger() zerolog.Logger {
	return log
}
