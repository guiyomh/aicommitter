package infrastructure

import (
	"github.com/google/wire"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
)

// ProviderSet est l'ensemble des providers pour l'infrastructure
var ProviderSet = wire.NewSet(
	NewCommandExecutor,
	wire.Bind(new(utils.Executor), new(*CommandExecutor)),
	LoadConfig,
)
