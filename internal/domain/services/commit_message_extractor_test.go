package services

import (
	"errors"
	"testing"

	"github.com/guiyomh/aicommitter/internal/domain/entities"
	"github.com/guiyomh/aicommitter/pkg/conventionalcommit"
	"github.com/stretchr/testify/assert"
)

func TestCommitMessageExtractor_Extract(t *testing.T) {
	tests := []struct {
		name        string
		response    AIResponse
		want        *entities.CommitMessage
		wantErr     bool
		expectedErr error
	}{
		{
			name: "réponse valide",
			response: `Some text before
---COMMIT_MESSAGE_START---
{
  "type": "feat",
  "scope": "auth",
  "description": "add google login",
  "body": "Implement Google OAuth integration",
  "footer": "Closes #123"
}
---COMMIT_MESSAGE_END---
Some text after`,
			want: &entities.CommitMessage{
				Type:        string(conventionalcommit.TypeFeat),
				Scope:       "auth",
				Description: "add google login",
				Body:        "Implement Google OAuth integration",
				Footer:      "Closes #123",
			},
			wantErr: false,
		},
		{
			name: "réponse minimale",
			response: `---COMMIT_MESSAGE_START---
{"type":"fix","description":"correct typo"}
---COMMIT_MESSAGE_END---`,
			want: &entities.CommitMessage{
				Type:        string(conventionalcommit.TypeFix),
				Description: "correct typo",
			},
			wantErr: false,
		},
		{
			name: "json invalide",
			response: `---COMMIT_MESSAGE_START---
{invalid json}
---COMMIT_MESSAGE_END---`,
			want:        nil,
			wantErr:     true,
			expectedErr: ErrInvalidJSON,
		},
		{
			name:        "sans délimiteur de début",
			response:    `{"type":"fix","description":"correct typo"}---COMMIT_MESSAGE_END---`,
			want:        nil,
			wantErr:     true,
			expectedErr: ErrMissingStartDelimiter,
		},
		{
			name:        "sans délimiteur de fin",
			response:    `---COMMIT_MESSAGE_START---{"type":"fix","description":"correct typo"}`,
			want:        nil,
			wantErr:     true,
			expectedErr: ErrMissingEndDelimiter,
		},
		{
			name:        "réponse vide",
			response:    "",
			want:        nil,
			wantErr:     true,
			expectedErr: ErrMissingStartDelimiter,
		},
		{
			name: "champs requis manquants - type",
			response: `---COMMIT_MESSAGE_START---
{"description":"test"}
---COMMIT_MESSAGE_END---`,
			want:    nil,
			wantErr: true,
			expectedErr: &ValidationError{
				Field:   "type",
				Message: "le champ est requis",
			},
		},
		{
			name: "champs requis manquants - description",
			response: `---COMMIT_MESSAGE_START---
{"type":"fix"}
---COMMIT_MESSAGE_END---`,
			want:    nil,
			wantErr: true,
			expectedErr: &ValidationError{
				Field:   "description",
				Message: "le champ est requis",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &MockParser{
				ParseFunc: func(message string) (*conventionalcommit.Commit, error) {
					if tt.wantErr {
						return nil, tt.expectedErr
					}
					return &conventionalcommit.Commit{
						Type:        conventionalcommit.Type(tt.want.Type),
						Scope:       tt.want.Scope,
						Description: tt.want.Description,
						Body:        tt.want.Body,
						Footer:      tt.want.Footer,
					}, nil
				},
			}
			extractor := NewCommitMessageExtractor(WithParser(parser))
			got, err := extractor.Extract(tt.response)
			if tt.wantErr {
				assert.Error(t, err)
				if tt.expectedErr != nil {
					var ve *ValidationError
					if errors.As(tt.expectedErr, &ve) {
						// Pour les erreurs de validation, on vérifie le type et les champs
						var gotVe *ValidationError
						if assert.ErrorAs(t, err, &gotVe) {
							assert.Equal(t, ve.Field, gotVe.Field)
							assert.Equal(t, ve.Message, gotVe.Message)
						}
					} else {
						// Pour les autres erreurs, on vérifie si l'erreur attendue est contenue dans l'erreur reçue
						assert.True(t, errors.Is(err, tt.expectedErr),
							"expected error %v, got %v", tt.expectedErr, err)
					}
				}
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
