package services

// ConventionalPromptGenerator implémente l'interface PromptGenerator
type ConventionalPromptGenerator struct {
	typeProvider CommitTypeProvider
	templateGen  PromptTemplateGenerator
}

// NewConventionalPromptGenerator crée une nouvelle instance
func NewConventionalPromptGenerator(
	typeProvider CommitTypeProvider,
	templateGen PromptTemplateGenerator,
) *ConventionalPromptGenerator {
	return &ConventionalPromptGenerator{
		typeProvider: typeProvider,
		templateGen:  templateGen,
	}
}

// GenerateCommitPrompt génère un prompt pour la création de messages de commit
func (g *ConventionalPromptGenerator) GenerateCommitPrompt(maxDiffSize int) string {
	commitTypes := g.typeProvider.GetCommitTypes()

	prompt := g.templateGen.GenerateTemplate(commitTypes)

	return prompt
}

// Static verification that ConventionalPromptGenerator implements PromptGenerator
var _ PromptGenerator = (*ConventionalPromptGenerator)(nil)
