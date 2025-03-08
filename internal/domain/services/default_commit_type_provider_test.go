package services

import (
	"testing"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/stretchr/testify/assert"
)

func TestNewDefaultCommitTypeProvider(t *testing.T) {
	// Arrange
	cfg := &config.Config{
		Git: struct {
			DefaultCommitStyle string                       `mapstructure:"default_commit_style"`
			CommitTypes        map[string]config.CommitType `mapstructure:"commit_types"`
			MaxDiffSize        int                          `mapstructure:"max_diff_size"`
		}{
			CommitTypes: map[string]config.CommitType{
				"feat": {
					Title:       "Features",
					Description: "A new feature",
					Emoji:       "✨",
				},
			},
		},
	}

	// Act
	provider := NewDefaultCommitTypeProvider(cfg)

	// Assert
	assert.NotNil(t, provider)
	assert.Equal(t, cfg, provider.config)
}

func TestGetCommitTypes(t *testing.T) {
	// Arrange
	commitTypes := map[string]config.CommitType{
		"feat": {
			Title:       "Features",
			Description: "A new feature",
			Emoji:       "✨",
		},
		"fix": {
			Title:       "Bug Fixes",
			Description: "A bug Fix",
			Emoji:       "🐛",
		},
	}
	cfg := &config.Config{
		Git: struct {
			DefaultCommitStyle string                       `mapstructure:"default_commit_style"`
			CommitTypes        map[string]config.CommitType `mapstructure:"commit_types"`
			MaxDiffSize        int                          `mapstructure:"max_diff_size"`
		}{
			CommitTypes: commitTypes,
		},
	}
	provider := NewDefaultCommitTypeProvider(cfg)

	// Act
	result := provider.GetCommitTypes()

	// Assert
	assert.Equal(t, commitTypes, result)
}

func TestGetTypeDescription(t *testing.T) {
	// Arrange
	commitTypes := map[string]config.CommitType{
		"feat": {
			Title:       "Features",
			Description: "A new feature",
			Emoji:       "✨",
		},
	}
	cfg := &config.Config{
		Git: struct {
			DefaultCommitStyle string                       `mapstructure:"default_commit_style"`
			CommitTypes        map[string]config.CommitType `mapstructure:"commit_types"`
			MaxDiffSize        int                          `mapstructure:"max_diff_size"`
		}{
			CommitTypes: commitTypes,
		},
	}
	provider := NewDefaultCommitTypeProvider(cfg)

	testCases := []struct {
		name           string
		commitType     string
		expectedResult string
	}{
		{
			name:           "Existing commit type",
			commitType:     "feat",
			expectedResult: "A new feature",
		},
		{
			name:           "Non-existing commit type",
			commitType:     "unknown",
			expectedResult: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := provider.GetTypeDescription(tc.commitType)

			// Assert
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestIsValidType(t *testing.T) {
	// Arrange
	commitTypes := map[string]config.CommitType{
		"feat": {
			Title:       "Features",
			Description: "A new feature",
			Emoji:       "✨",
		},
	}
	cfg := &config.Config{
		Git: struct {
			DefaultCommitStyle string                       `mapstructure:"default_commit_style"`
			CommitTypes        map[string]config.CommitType `mapstructure:"commit_types"`
			MaxDiffSize        int                          `mapstructure:"max_diff_size"`
		}{
			CommitTypes: commitTypes,
		},
	}
	provider := NewDefaultCommitTypeProvider(cfg)

	testCases := []struct {
		name           string
		commitType     string
		expectedResult bool
	}{
		{
			name:           "Valid commit type",
			commitType:     "feat",
			expectedResult: true,
		},
		{
			name:           "Invalid commit type",
			commitType:     "unknown",
			expectedResult: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result := provider.IsValidType(tc.commitType)

			// Assert
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
