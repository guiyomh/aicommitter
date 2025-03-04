package services

// AIProvider defines the interface for AI providers could be used to generate commit messages
type AIProvider interface {
	// GenerateCommitMessage generates a commit message based on the given input
	GenerateCommitMessage(diff string, changedFiles []string) (string, error)
}
