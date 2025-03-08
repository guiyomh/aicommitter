package services

import (
	"errors"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/guiyomh/aicommitter/pkg/conventionalcommit"
)

const (
	startDelimiter = "---COMMIT_MESSAGE_START---"
	endDelimiter   = "---COMMIT_MESSAGE_END---"
)

var (
	ErrInvalidJSON           = errors.New("json invalide")
	ErrMissingStartDelimiter = errors.New("délimiteur de début manquant")
	ErrMissingEndDelimiter   = errors.New("délimiteur de fin manquant")
)

// ValidationError représente une erreur de validation des champs
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// CommitMessageExtractor extrait un message de commit à partir d'une réponse
type CommitMessageExtractor struct {
	parser conventionalcommit.Parser
}

// Option est une fonction qui configure le CommitMessageExtractor
type Option func(*CommitMessageExtractor)

// WithParser configure le parser à utiliser
func WithParser(parser conventionalcommit.Parser) Option {
	return func(e *CommitMessageExtractor) {
		e.parser = parser
	}
}

// NewCommitMessageExtractor crée une nouvelle instance de CommitMessageExtractor
func NewCommitMessageExtractor(opts ...Option) *CommitMessageExtractor {
	e := &CommitMessageExtractor{
		parser: conventionalcommit.NewParser(),
	}
	for _, opt := range opts {
		opt(e)
	}
	return e
}

// Extract extrait un message de commit à partir d'une réponse
func (e *CommitMessageExtractor) Extract(response AIResponse) (*entities.CommitMessage, error) {
	rawResponse := string(response)

	// Validation du type de commit via le parser
	parsedCommit, err := e.parser.Parse(strings.TrimSpace(rawResponse))
	if err != nil {
		if conventionalcommit.IsValidationError(err) {
			return nil, &ValidationError{Field: "type", Message: "type de commit invalide"}
		}
		return nil, err
	}

	commit := entities.CommitMessage{
		Type:        string(parsedCommit.Type),
		Scope:       parsedCommit.Scope,
		Description: parsedCommit.Description,
		Body:        parsedCommit.Body,
		Footer:      parsedCommit.Footer,
		Breaking:    parsedCommit.Breaking,
	}

	return &commit, nil
}
