package domain

import (
	"github.com/google/wire"
	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/guiyomh/aicommitter/internal/domain/usecases"
)

// ProviderSet is the set of providers for wire to inject the DiffService
var ProviderSet = wire.NewSet(
	services.NewDiffService,
	usecases.NewGenerateCommitMessage,
	services.NewDefaultPromptGenerator,
	wire.Bind(new(services.PromptGenerator), new(*services.DefaultPromptGenerator)),
)
