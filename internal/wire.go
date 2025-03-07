//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/google/wire"
	"github.com/guiyomh/aicommitter/internal/cli/commands"
	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/infrastructure"
	"github.com/spf13/cobra"
)

// InitializeCommitCommand initializes the commit command with its dependencies
func CommitCommand() (*cobra.Command, error) {
	wire.Build(
		infrastructure.ProviderSet,
		domain.ProviderSet,
		commands.CommitCommand,
	)
	return nil, nil
}

// InitializeInitConfigCommand initializes the init-config command with its dependencies
func InitConfigCommand() *cobra.Command {
	wire.Build(
		commands.InitConfigCommand,
	)
	return nil
}
