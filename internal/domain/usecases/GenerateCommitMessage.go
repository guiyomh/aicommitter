package usecases

import (
	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/services"
)

// GenerateCommitMessage is a use case to generate a commit message with the help of an AI
type GenerateCommitMessage struct {
	diffService *services.DiffService
	aiProvider  services.AIProvider
	config      *config.Config
}

// NewGenerateCommitMessage creates a new instance of the use case
func NewGenerateCommitMessage(
	diffService *services.DiffService,
	aiProvider services.AIProvider,
	config *config.Config,
) *GenerateCommitMessage {
	return &GenerateCommitMessage{
		diffService: diffService,
		aiProvider:  aiProvider,
		config:      config,
	}
}

func (uc *GenerateCommitMessage) Execute(staged bool) (string, error) {
	var diff string
	var err error

	if staged {
		diff, err = uc.diffService.GetStagedDiff()
	} else {
		diff, err = uc.diffService.GetDiff()
	}

	if err != nil {
		return "", err
	}

	changedFiles, err := uc.diffService.GetFilesChanged()
	if err != nil {
		return "", err
	}

	commitMessage, err := uc.aiProvider.GenerateCommitMessage(diff, changedFiles)

	if err != nil {
		return "", err
	}

	return commitMessage, nil
}
