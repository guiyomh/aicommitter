package services

import (
	"encoding/json"
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
func (e *CommitMessageExtractor) Extract(response string) (*entities.CommitMessage, error) {
	startIndex := strings.Index(response, startDelimiter)
	if startIndex == -1 {
		return nil, ErrMissingStartDelimiter
	}

	endIndex := strings.Index(response, endDelimiter)
	if endIndex == -1 {
		return nil, ErrMissingEndDelimiter
	}

	jsonStr := strings.TrimSpace(response[startIndex+len(startDelimiter) : endIndex])
	var commit entities.CommitMessage
	if err := json.Unmarshal([]byte(jsonStr), &commit); err != nil {
		return nil, ErrInvalidJSON
	}

	// Validation des champs requis
	if commit.Type == "" {
		return nil, &ValidationError{Field: "type", Message: "le champ est requis"}
	}
	if commit.Description == "" {
		return nil, &ValidationError{Field: "description", Message: "le champ est requis"}
	}

	// Validation du type de commit via le parser
	_, err := e.parser.Parse(commit.String())
	if err != nil {
		if conventionalcommit.IsValidationError(err) {
			return nil, &ValidationError{Field: "type", Message: "type de commit invalide"}
		}
		return nil, err
	}

	return &commit, nil
}
