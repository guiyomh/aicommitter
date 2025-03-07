package infrastructure

import (
	"fmt"
	"os"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ZerologLogger is an implementation of the Logger interface using zerolog
type ZerologLogger struct{}

// NewLogger configures and creates a new instance of ZerologLogger
func NewLogger(config *config.Config) *ZerologLogger {
	// Configurer zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	// Configurer le niveau de log
	logLevel, err := zerolog.ParseLevel(config.Log.Level)
	if err != nil {
		// Par défaut, utiliser info en cas d'erreur
		logLevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(logLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	return &ZerologLogger{}
}

// Debug records a debug level message
func (*ZerologLogger) Debug(msg string, args ...interface{}) {
	log.Debug().Msg(formatMessage(msg, args...))
}

// Info records an info level message
func (*ZerologLogger) Info(msg string, args ...interface{}) {
	log.Info().Msg(formatMessage(msg, args...))
}

// Warn records a warn level message
func (*ZerologLogger) Warn(msg string, args ...interface{}) {
	log.Warn().Msg(formatMessage(msg, args...))
}

// Error records an error level message
func (*ZerologLogger) Error(msg string, args ...interface{}) {
	log.Error().Msg(formatMessage(msg, args...))
}

// Fatal records a fatal level message, then terminates the program
func (*ZerologLogger) Fatal(msg string, args ...interface{}) {
	log.Fatal().Msg(formatMessage(msg, args...))
}

// formatMessage formats the message if arguments are supplied
func formatMessage(msg string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}

// Check static that ZerologLogger implements the Logger interface
var _ utils.Logger = (*ZerologLogger)(nil)
