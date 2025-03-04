package services

import (
	"fmt"
	"strings"
)

// DefaultPromptGenerator implémente l'interface PromptGenerator avec une logique par défaut
type DefaultPromptGenerator struct{}

// NewDefaultPromptGenerator crée une nouvelle instance de DefaultPromptGenerator
func NewDefaultPromptGenerator() *DefaultPromptGenerator {
	return &DefaultPromptGenerator{}
}

// GenerateCommitPrompt génère un prompt pour la création de messages de commit
func (g *DefaultPromptGenerator) GenerateCommitPrompt(diff string, changedFiles []string, maxDiffSize int) string {
	if len(diff) > maxDiffSize {
		diff = diff[:maxDiffSize] + "\n[diff truncated...]"
	}

	fileStr := strings.Join(changedFiles, ",")

	return fmt.Sprintf(
		"Generates a commit message in conventional format for the following changes.\n\n"+
			"Files modified: %s\n\n"+
			"Diff:\n%s\n\n"+
			"Format: <type>(<scope>): <description>\n\n"+
			"Where <type> is one of the following: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert\n"+
			"<scope> is optional and represents the affected part of the code\n "+
			"<description> is a short description of the modifications\n\n",
		fileStr,
		diff,
	)
}

// Vérification statique que DefaultPromptGenerator implémente bien l'interface PromptGenerator
var _ PromptGenerator = (*DefaultPromptGenerator)(nil)
