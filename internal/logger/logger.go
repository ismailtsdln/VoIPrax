package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger represents the system logger
type Logger struct {
	zerolog.Logger
}

// New initializes and returns a new Logger instance
func New(level string, pretty bool) *Logger {
	var output zerolog.ConsoleWriter
	if pretty {
		output = zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	}

	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}

	var l zerolog.Logger
	if pretty {
		l = zerolog.New(output).With().Timestamp().Logger().Level(lvl)
	} else {
		l = zerolog.New(os.Stderr).With().Timestamp().Logger().Level(lvl)
	}

	return &Logger{l}
}

// Global initialization for standard log package compatibility if needed
func InitGlobal(level string, pretty bool) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)

	if pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}
}
