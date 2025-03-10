package services

import (
	"errors"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/models"
	"github.com/guiyomh/aicommitter/pkg/conventionalcommit"
)

const (
	startDelimiter = "---COMMIT_MESSAGE_START---"
	endDelimiter   = "---COMMIT_MESSAGE_END---"
	// DefaultCommitType is used when the commit type is invalid or not provided
	DefaultCommitType = "chore"
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

// Extract extracts a commit message from an AI response
func (e *CommitMessageExtractor) Extract(response AIResponse) (*models.CommitMessage, error) {
	rawResponse := string(response)
	// Try to parse the full commit message
	parsedCommit, err := e.parser.Parse(strings.TrimSpace(rawResponse))
	if err == nil {
		return &models.CommitMessage{
			Type:        string(parsedCommit.Type),
			Scope:       parsedCommit.Scope,
			Description: parsedCommit.Description,
			Body:        parsedCommit.Body,
			Footer:      parsedCommit.Footer,
			Breaking:    parsedCommit.Breaking,
		}, nil
	}

	// If parsing failed, check if we have partial information
	if conventionalcommit.IsValidationError(err) {
		validationErr := conventionalcommit.GetValidationError(err)
		if validationErr.PartialCommit != nil {
			// Create a commit message with the partial information
			commitMsg := &models.CommitMessage{
				Type:        DefaultCommitType, // Use default type unless we have a valid one
				Description: validationErr.PartialCommit.Description,
				Scope:       validationErr.PartialCommit.Scope,
				Body:        validationErr.PartialCommit.Body,
				Footer:      validationErr.PartialCommit.Footer,
				Breaking:    validationErr.PartialCommit.Breaking,
			}

			// If we have a valid type from partial parsing, use it
			if validationErr.PartialCommit.Type != "" {
				commitMsg.Type = string(validationErr.PartialCommit.Type)
			}

			return commitMsg, nil
		}
	}

	// If we have no partial information, return the error
	return nil, err
}

func isValidType(t string) bool {
	validTypes := []string{
		"feat", "fix", "docs", "style", "refactor",
		"perf", "test", "build", "ci", "chore", "revert",
	}
	t = strings.TrimSpace(strings.ToLower(t))
	for _, vt := range validTypes {
		if t == vt {
			return true
		}
	}
	return false
}
