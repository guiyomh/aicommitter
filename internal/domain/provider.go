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
	services.NewConventionalPromptGenerator,
	services.NewConventionalTemplateGenerator,
	services.NewDefaultCommitTypeProvider,
	services.NewDefaultDiffFormatter,
	wire.Bind(new(services.PromptGenerator), new(*services.ConventionalPromptGenerator)),
	wire.Bind(new(services.CommitTypeProvider), new(*services.DefaultCommitTypeProvider)),
	wire.Bind(new(services.PromptTemplateGenerator), new(*services.ConventionTemplateGenerator)),
	wire.Bind(new(services.DiffFormatter), new(*services.DefaultDiffFormatter)),
)
