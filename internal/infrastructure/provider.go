package infrastructure

import (
	"github.com/google/wire"
	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
	"github.com/guiyomh/aicommitter/internal/infrastructure/ai"
)

func NewOllamaProvider(config *config.Config) (*ai.OllamaProvider, error) {
	return ai.NewOllamaProvider(
		config.AI.BaseURL,
		config.AI.Model,
		config.Git.MaxDiffSize,
	)
}

// ProviderSet est l'ensemble des providers pour l'infrastructure
var ProviderSet = wire.NewSet(
	NewCommandExecutor,
	wire.Bind(new(utils.Executor), new(*CommandExecutor)),
	LoadConfig,
	NewOllamaProvider,
	wire.Bind(new(services.AIProvider), new(*ai.OllamaProvider)),
)
