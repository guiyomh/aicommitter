package infrastructure

import (
	"fmt"
	"os"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ZerologLogger est une implémentation de l'interface Logger utilisant zerolog
type ZerologLogger struct{}

// NewLogger configure et crée une nouvelle instance de ZerologLogger
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

// Debug enregistre un message de niveau debug
func (l *ZerologLogger) Debug(msg string, args ...interface{}) {
	log.Debug().Msg(formatMessage(msg, args...))
}

// Info enregistre un message de niveau info
func (l *ZerologLogger) Info(msg string, args ...interface{}) {
	log.Info().Msg(formatMessage(msg, args...))
}

// Warn enregistre un message de niveau warn
func (l *ZerologLogger) Warn(msg string, args ...interface{}) {
	log.Warn().Msg(formatMessage(msg, args...))
}

// Error enregistre un message de niveau error
func (l *ZerologLogger) Error(msg string, args ...interface{}) {
	log.Error().Msg(formatMessage(msg, args...))
}

// Fatal enregistre un message de niveau fatal puis termine le programme
func (l *ZerologLogger) Fatal(msg string, args ...interface{}) {
	log.Fatal().Msg(formatMessage(msg, args...))
}

// formatMessage formate le message si des arguments sont fournis
func formatMessage(msg string, args ...interface{}) string {
	if len(args) > 0 {
		return fmt.Sprintf(msg, args...)
	}
	return msg
}

// Vérification statique que ZerologLogger implémente bien l'interface Logger
var _ utils.Logger = (*ZerologLogger)(nil)
