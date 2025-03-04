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

// InitializeCommitCommand initialise la commande commit avec ses dépendances
func CommitCommand() *cobra.Command {
	wire.Build(
		infrastructure.ProviderSet,
		domain.ProviderSet,
		commands.CommitCommand,
	)
	return nil
}
