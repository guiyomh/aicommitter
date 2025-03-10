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
	return strings.Join([]string{
		"You are a helpful assistant specializing in writing clear and informative Git commit messages using the Conventional Commits style.",
		"Based on the given Git diff and changed files, generate exactly one conventional commit message following these guidelines:",
		"",
		"1. Message Language: English",
		"2. Format: follow the Conventional Commits format:",
		"    <type>(<optional scope>): <description>",
		"    ",
		"    <optional body>",
		"    ",
		"    <optional footer>",
		"    ",
		"3. Types: use one of the following types:",
		strings.Join(
			lo.Map(
				lo.Keys(commitTypes),
				func(key string, _ int) string {
					return fmt.Sprintf("   - %s: %s", key, commitTypes[key].Description)
				},
			),
			"\n",
		),
		"4. Guidelines for writing commit messages:",
		"   - Be specific about what changes were made",
		"   - Use imperative mood (\"add feature\" not \"added feature\")",
		"   - Keep subject line under 50 characters",
		"   - Do not end the subject line with a period",
		"   - Use the body to explain what and why vs. how",
		"5. Focus on:",
		"   - What problem this commit solves",
		"   - Why this change was necessary",
		"   - Any important technical details",
		"6. Exclude anything unnecessary such as translation or implementation details.",
	}, "\n")
}
