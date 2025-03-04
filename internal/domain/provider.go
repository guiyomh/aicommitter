package domain

import (
	"github.com/google/wire"
	"github.com/guiyomh/aicommitter/internal/domain/services"
)

// ProviderSet is the set of providers for wire to inject the DiffService
var ProviderSet = wire.NewSet(services.NewDiffService)
