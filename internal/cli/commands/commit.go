package commands

import (
	"context"
	"fmt"

	"github.com/guiyomh/aicommitter/internal/cli/ui"
	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/usecases"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
	"github.com/spf13/cobra"
)

func CommitCommand(
	generateCommitMessage *usecases.GenerateCommitMessage,
	cfg *config.Config,
	log utils.Logger,
	executor utils.Executor,
) *cobra.Command {
	var staged bool

	commit := &cobra.Command{
		Use:   "commit",
		Short: "Commit a model",
		RunE: func(_ *cobra.Command, _ []string) error {
			// Generate initial commit message
			commitMessage, err := generateCommitMessage.Execute(staged)
			if err != nil {
				return fmt.Errorf("error while generating commit message: %w", err)
			}
			log.Info("Commit message: %v", commitMessage)

			// Create and run the TUI editor
			editor := ui.NewCommitEditor(commitMessage, cfg.Git.CommitTypes)
			finalMessage, confirm, err := editor.Run()

			if err != nil {
				return fmt.Errorf("error while editing commit message: %w", err)
			}

			// Use the final message
			fmt.Println("Message de commit final :")
			fmt.Println(finalMessage.String())
			if confirm {
				executor.Execute(context.Background(), "git", "commit", "-m", finalMessage.String())
				log.Info("Commit message confirmed")
			} else {
				fmt.Println("Commit message not confirmed")
			}

			return nil
		},
	}

	commit.Flags().BoolVarP(&staged, "staged", "s", false, "Uses only changes already added (staging)")

	return commit
}
