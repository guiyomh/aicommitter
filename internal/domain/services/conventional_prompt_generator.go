package services

import (
	"fmt"

	"github.com/samber/lo"
)

// ConventionalPromptGenerator implémente l'interface PromptGenerator
type ConventionalPromptGenerator struct {
	typeProvider  CommitTypeProvider
	templateGen   PromptTemplateGenerator
	diffFormatter DiffFormatter
}

// NewConventionalPromptGenerator crée une nouvelle instance
func NewConventionalPromptGenerator(
	typeProvider CommitTypeProvider,
	templateGen PromptTemplateGenerator,
	diffFormatter DiffFormatter,
) *ConventionalPromptGenerator {
	return &ConventionalPromptGenerator{
		typeProvider:  typeProvider,
		templateGen:   templateGen,
		diffFormatter: diffFormatter,
	}
}

// GenerateCommitPrompt génère un prompt pour la création de messages de commit
func (g *ConventionalPromptGenerator) GenerateCommitPrompt(diff string, changedFiles []string, maxDiffSize int) string {
	commitTypes := lo.Keys(g.typeProvider.GetCommitTypes())
	formattedDiff := g.diffFormatter.FormatDiff(diff, maxDiffSize)
	fileStr := g.diffFormatter.GetFileList(changedFiles)

	template := g.templateGen.GenerateTemplate(commitTypes)

	return fmt.Sprintf(template, formattedDiff, fileStr)
}

// Static verification that ConventionalPromptGenerator implements PromptGenerator
var _ PromptGenerator = (*ConventionalPromptGenerator)(nil)
