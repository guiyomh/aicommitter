package services

import (
	"github.com/guiyomh/aicommitter/internal/domain/config"
)

// PromptGenerator defines the interface for prompt generators
type PromptGenerator interface {
	// GenerateCommitPrompt generates a commit prompt based on the given input
	GenerateCommitPrompt(diff string, changedFiles []string, maxDiffSize int) string
}

// CommitTypeProvider defines the interface for commit type providers
type CommitTypeProvider interface {
	GetCommitTypes() map[string]config.CommitType
	GetTypeDescription(commitType string) string
	IsValidType(commitType string) bool
}

// PromptTemplateGenerator defines the interface for prompt template generators
type PromptTemplateGenerator interface {
	GenerateTemplate(commitTypes map[string]config.CommitType) string
}

// DiffFormatter defines the interface for diff formatters
type DiffFormatter interface {
	FormatDiff(diff string, maxSize int) string
	GetFileList(files []string) string
}

// AIProvider defines the interface for AI providers
type AIProvider interface {
	// GenerateCommitMessage generates a commit message based on the given input
	GenerateCommitMessage(diff string, changedFiles []string) (string, error)
}
