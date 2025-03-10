package ui

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/models"
)

type CommitEditor struct {
	form          *huh.Form
	commitMessage *models.CommitMessage
	commitConfirm *bool
}

func NewCommitEditor(commitMessage *models.CommitMessage, commitTypes map[string]config.CommitType) *CommitEditor {

	commitTypeOptions := make([]huh.Option[string], 0, len(commitTypes))
	resultCommitMessage := *commitMessage // Create a copy by dereferencing and copying the value
	for key, commitType := range commitTypes {
		commitTypeOptions = append(commitTypeOptions, huh.NewOption(
			fmt.Sprintf("%s %s: %s %s", commitType.Emoji, key, commitType.Title, commitType.Description),
			key,
		))
	}
	confirmConfirm := false

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Commit type").
				Options(commitTypeOptions...).
				Value(&resultCommitMessage.Type),
			huh.NewInput().
				Title("Commit scope").
				Placeholder("Write your commit scope here...").
				Value(&resultCommitMessage.Scope),
			huh.NewInput().
				Title("Commit message").
				Placeholder("Write your commit message here...").
				Value(&resultCommitMessage.Description),
			huh.NewText().
				Title("Commit body").
				Placeholder("Write your commit body here...").
				Value(&resultCommitMessage.Body),
			huh.NewText().
				Title("Commit footer").
				Placeholder("Write your commit footer here...").
				Value(&resultCommitMessage.Footer),
			huh.NewConfirm().
				Title("Is this a breaking change?").
				Value(&resultCommitMessage.Breaking),
		),
		huh.NewGroup(
			huh.NewNote().
				Title("Preview of your commit message").
				Description(resultCommitMessage.String()),
			huh.NewConfirm().
				Title("Do you confirm this commit message?").
				Value(&confirmConfirm),
		),
	)

	return &CommitEditor{
		form:          form,
		commitMessage: &resultCommitMessage,
		commitConfirm: &confirmConfirm,
	}
}

func (e *CommitEditor) Run() (*models.CommitMessage, bool, error) {
	err := e.form.Run()
	if err != nil {
		return nil, false, err
	}
	return e.commitMessage, *e.commitConfirm, nil
}
