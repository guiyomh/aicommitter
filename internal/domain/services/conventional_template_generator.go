package services

import (
	"fmt"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/samber/lo"
)

// ConventionTemplateGenerator implémente PromptTemplateGenerator
type ConventionTemplateGenerator struct{}

func NewConventionalTemplateGenerator() *ConventionTemplateGenerator {
	return &ConventionTemplateGenerator{}
}

func (*ConventionTemplateGenerator) GenerateTemplate(commitTypes map[string]config.CommitType) string {
	return `You are a commit message generator.
Your ONLY task is to analyze the Git diff and generate a conventional commit message.

IMPORTANT: Generate ONLY plain ASCII characters. NO emojis, NO special characters, NO escape sequences.

# Input Context

## Git Diff:
` + "```" + `
%s
` + "```" + `

## Changed Files:
` + "```" + `
%s
` + "```" + `

# Output Format
YOU MUST OUTPUT EXACTLY THIS FORMAT, WITH NO VARIATIONS:

` + "```\n" + startDelimiter + `
<type>[(scope)][!]: <description>

[body]

[footer]
` + endDelimiter + "\n```" + `

# Format Rules
- Use ONLY plain ASCII characters
- NO emojis or special characters
- NO extra spaces or tabs
- Add a "!" after the scope (before the colon) to indicate a breaking change

# Available Commit Types:
` + strings.Join(
		lo.Map(
			lo.Keys(commitTypes),
			func(key string, _ int) string {
				return fmt.Sprintf("- %s: %s", key, commitTypes[key].Description)
			},
		),
		"\n",
	) + `

# Content Rules
1. type: Must be one of the types listed above
2. scope: Optional. Use parentheses when present. Omit completely if not needed
3. description: < 50 chars, imperative mood, lowercase start, no period
4. body: Optional. Use to explain the motivation for the change and contrast it with previous behavior
5. footer: Optional. Use for referencing issues (e.g., "Fixes #123")
6. breaking: Add "!" after scope (or type if no scope) to indicate breaking changes

# Critical Instructions
- OUTPUT ONLY THE COMMIT MESSAGE BETWEEN DELIMITERS
- NO TEXT BEFORE OR AFTER DELIMITERS
- USE EXACT DELIMITER SPELLING
- ENSURE PROPER CONVENTIONAL COMMIT FORMAT
- USE ONLY ASCII CHARACTERS

Example of exact expected format:
` + "```\n" + startDelimiter + `
feat(auth)!: add google oauth login

Implement Google OAuth for user authentication
Add sign-in button to login page

Fixes #123
` + endDelimiter + "\n```" + `
`
}
