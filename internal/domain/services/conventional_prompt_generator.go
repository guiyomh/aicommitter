package services

import (
	"fmt"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/samber/lo"
)

// ConventionalPromptGenerator implémente l'interface PromptGenerator avec une logique par défaut
type ConventionalPromptGenerator struct {
	commitTypes map[string]config.CommitType
}

// NewDefaultPromptGenerator crée une nouvelle instance de DefaultPromptGenerator
func NewDefaultPromptGenerator(cfg *config.Config) *ConventionalPromptGenerator {
	return &ConventionalPromptGenerator{
		commitTypes: cfg.Git.CommitTypes,
	}
}

// GenerateCommitPrompt génère un prompt pour la création de messages de commit
func (g *ConventionalPromptGenerator) GenerateCommitPrompt(diff string, changedFiles []string, maxDiffSize int) string {
	if len(diff) > maxDiffSize {
		diff = diff[:maxDiffSize] + "\n[diff truncated...]"
	}

	fileStr := strings.Join(changedFiles, ",")

	return `You are an assistant specialized in generating commit messages following the Conventional Commits format.

# Objective
Generate a descriptive, clear and concise commit message that strictly follows the Conventional Commits specification.

# Commit Context

- Git Diff:

` + "```\n" + diff + "\n```" + `

- Impacted files: ` + fileStr + `

# Constraints and Instructions

## Commit Format

The message MUST follow the format:

` + "```" + `
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
` + "```" + `

## Generation Rules

1. Commit Type Choice

- Use one of the standard types:

  - ` + strings.Join(
		lo.Map(
			lo.Keys(g.commitTypes),
			func(key string, _ int) string {
				return fmt.Sprintf("%s: %s", key, g.commitTypes[key].Description)
			},
		),
		"\n  - ",
	) + `

2. Scope (Optional)

- Indicate the main affected component/module/file
- Limited to technical or functional scope

3. Description

- Short (< 50 characters)
- Imperative mood, like "Add/Fix/Modify..."
- Start with lowercase letter
- No trailing period

4. Message Body (Optional)

- More detailed explanation if needed
- Motivation for the change
- Differences from previous state

5. Footer(s) (Optional)

- Issue references (e.g. "Fixes #123")
- Meta information
` +

		// # Paramètres Spécifiques Fournis

		// - Type forcé : {type_force}
		// - Scope forcé : {scope_force}
		// - Numéro d'issue forcé : {issue_force}
		// - Langue du message : {langue}

		`
# Additional Instructions

- Be concise and precise
- Use the diff context to understand the changes
- If no type is obvious, choose the most appropriate one
- When in doubt, prefer refactor or chore

# Response Format
Reply ONLY with the commit message, without any additional text.
The message MUST be easily parsable, so:

- Use clear delimiters like ---COMMIT_MESSAGE_START--- and ---COMMIT_MESSAGE_END---
- Include parsable metadata

Example response format:

` + "```" + `
---COMMIT_MESSAGE_START---
{
  "type": "feat",
  "scope": "authentication",
  "description": "add login with google oauth",
  "body": "Implement Google OAuth integration for user authentication\n` +
		`Adds support for Google sign-in on the login page",
  "footer": "Fixes #456"
}
---COMMIT_MESSAGE_END---
` + "```\n"
}

// Vérification statique que DefaultPromptGenerator implémente bien l'interface PromptGenerator
var _ PromptGenerator = (*ConventionalPromptGenerator)(nil)
