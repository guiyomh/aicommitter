package services

// PromptGenerator est responsable de la génération de prompts pour les modèles IA
type PromptGenerator interface {
	// GenerateCommitPrompt génère un prompt pour la création de messages de commit
	GenerateCommitPrompt(diff string, changedFiles []string, maxDiffSize int) string
}
