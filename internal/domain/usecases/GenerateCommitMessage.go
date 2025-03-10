package usecases

import (
	"fmt"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/models"
	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
)

// GenerateCommitMessage is a use case to generate a commit message with the help of an AI
type GenerateCommitMessage struct {
	diffService            *services.DiffService
	aiProvider             services.AIProvider
	config                 *config.Config
	log                    utils.Logger
	commitMessageExtractor *services.CommitMessageExtractor
}

// NewGenerateCommitMessage creates a new instance of the use case
func NewGenerateCommitMessage(
	diffService *services.DiffService,
	aiProvider services.AIProvider,
	config *config.Config,
	log utils.Logger,
	commitMessageExtractor *services.CommitMessageExtractor,
) *GenerateCommitMessage {
	return &GenerateCommitMessage{
		diffService:            diffService,
		aiProvider:             aiProvider,
		config:                 config,
		log:                    log,
		commitMessageExtractor: commitMessageExtractor,
	}
}

func (uc *GenerateCommitMessage) Execute(staged bool) (*models.CommitMessage, error) {
	var diff string
	var err error

	if staged {
		diff, err = uc.diffService.GetStagedDiff()
	} else {
		diff, err = uc.diffService.GetDiff()
	}

	uc.log.Debug("Diff: %s", diff)
	if err != nil {
		return nil, err
	}

	if len(diff) == 0 {
		return nil, fmt.Errorf("no changes to commit")
	}

	changedFiles, err := uc.diffService.GetFilesChanged()
	if err != nil {
		return nil, err
	}

	iaResponse, err := uc.aiProvider.GenerateCommitMessage(diff, changedFiles)
	if err != nil {
		return nil, err
	}

	return uc.commitMessageExtractor.Extract(iaResponse)
}
