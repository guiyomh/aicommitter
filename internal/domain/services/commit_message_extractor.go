package services

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
)

const (
	startDelimiter = "---COMMIT_MESSAGE_START---"
	endDelimiter   = "---COMMIT_MESSAGE_END---"
)

// Erreurs personnalisées pour l'extracteur
var (
	ErrMissingStartDelimiter = fmt.Errorf("délimiteur de début non trouvé dans la réponse")
	ErrMissingEndDelimiter   = fmt.Errorf("délimiteur de fin non trouvé dans la réponse")
	ErrInvalidJSON           = fmt.Errorf("le contenu extrait n'est pas un JSON valide")
)

// ValidationError représente une erreur de validation du message de commit
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation du champ '%s': %s", e.Field, e.Message)
}

// NewValidationError crée une nouvelle erreur de validation
func NewValidationError(field, message string) *ValidationError {
	return &ValidationError{
		Field:   field,
		Message: message,
	}
}

// CommitMessageExtractor extrait et convertit le message de commit de la réponse de l'IA
type CommitMessageExtractor struct{}

// NewCommitMessageExtractor crée une nouvelle instance de CommitMessageExtractor
func NewCommitMessageExtractor() *CommitMessageExtractor {
	return &CommitMessageExtractor{}
}

// validateCommitMessage vérifie que les champs requis sont présents
func (e *CommitMessageExtractor) validateCommitMessage(msg *entities.CommitMessage) error {
	if msg.Type == "" {
		return NewValidationError("type", "le champ est requis")
	}
	if msg.Description == "" {
		return NewValidationError("description", "le champ est requis")
	}
	return nil
}

// Extract extrait et convertit le message de commit de la réponse de l'IA en entité CommitMessage
func (e *CommitMessageExtractor) Extract(aiResponse string) (*entities.CommitMessage, error) {
	// Trouver l'index du début du message
	startIndex := strings.Index(aiResponse, startDelimiter)
	if startIndex == -1 {
		return nil, ErrMissingStartDelimiter
	}
	startIndex += len(startDelimiter)

	// Trouver l'index de fin du message
	endIndex := strings.Index(aiResponse, endDelimiter)
	if endIndex == -1 {
		return nil, ErrMissingEndDelimiter
	}

	// Extraire le contenu entre les délimiteurs
	content := strings.TrimSpace(aiResponse[startIndex:endIndex])

	// Convertir en entité CommitMessage
	var commitMessage entities.CommitMessage
	if err := json.Unmarshal([]byte(content), &commitMessage); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrInvalidJSON, err)
	}

	// Valider les champs requis
	if err := e.validateCommitMessage(&commitMessage); err != nil {
		return nil, err
	}

	return &commitMessage, nil
}
