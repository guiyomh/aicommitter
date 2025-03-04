package commands

import (
	"fmt"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/usecases"
	"github.com/spf13/cobra"
)

func CommitCommand(
	generateCommitMessage *usecases.GenerateCommitMessage,
	cfg *config.Config,
) *cobra.Command {

	var staged bool

	commit := &cobra.Command{
		Use:   "commit",
		Short: "Commit a model",
		RunE: func(_ *cobra.Command, _ []string) error {

			commitMessage, err := generateCommitMessage.Execute(staged)

			if err != nil {
				return fmt.Errorf("error while generating commit message: %w", err)
			}

			fmt.Println("Message de commit généré :")
			fmt.Println(commitMessage)

			return nil
		},
	}

	commit.Flags().BoolVarP(&staged, "staged", "s", true, "Uses only changes already added (staging)")

	return commit
}
