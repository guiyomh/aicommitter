package commands

import "github.com/google/wire"

// ProviderSet est l'ensemble des providers pour le package cli
var ProviderSet = wire.NewSet(CommitCommand)
