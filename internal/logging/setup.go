package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New(level zerolog.Level) *zerolog.Logger {
	writer := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	}

	logger := zerolog.New(writer).Level(level).With().Timestamp().Logger()
	return &logger
}
