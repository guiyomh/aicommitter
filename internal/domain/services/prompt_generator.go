package services

// PromptGenerator is responsible for generating prompts for IA models
type PromptGenerator interface {
	// GenerateCommitPrompt génère un prompt pour la création de messages de commit
	GenerateCommitPrompt(diff string, changedFiles []string, maxDiffSize int) string
}
