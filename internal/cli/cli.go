package cli

import (
	"fmt"
	"os"

	"github.com/guiyomh/aicommitter/internal"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	root := &cobra.Command{
		Use:   "aicommiter",
		Short: "AI-powered commit message generator",
	}

	// Add commit subcommand
	commitCmd, err := internal.CommitCommand()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create commit command: %v\n", err)
		os.Exit(1)
	}
	root.AddCommand(commitCmd)
	// Add init-config subcommand
	initConfigCmd := internal.InitConfigCommand()
	root.AddCommand(initConfigCmd)

	return root
}
