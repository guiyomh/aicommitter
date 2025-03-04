package cli

import (
	"github.com/guiyomh/aicommitter/internal"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	root := &cobra.Command{
		Use:   "aicommiter",
		Short: "AI-powered commit message generator",
	}
	commitCmd := internal.CommitCommand()
	root.AddCommand(commitCmd)

	return root
}
