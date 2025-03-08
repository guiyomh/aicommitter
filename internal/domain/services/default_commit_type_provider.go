package services

import "github.com/guiyomh/aicommitter/internal/domain/config"

// DefaultCommitTypeProvider implements CommitTypeProvider
type DefaultCommitTypeProvider struct {
	config *config.Config
}

func NewDefaultCommitTypeProvider(cfg *config.Config) *DefaultCommitTypeProvider {
	return &DefaultCommitTypeProvider{config: cfg}
}

func (p *DefaultCommitTypeProvider) GetCommitTypes() map[string]config.CommitType {
	return p.config.Git.CommitTypes
}

func (p *DefaultCommitTypeProvider) GetTypeDescription(commitType string) string {
	if ct, exists := p.config.Git.CommitTypes[commitType]; exists {
		return ct.Description
	}
	return ""
}

func (p *DefaultCommitTypeProvider) IsValidType(commitType string) bool {
	_, exists := p.config.Git.CommitTypes[commitType]
	return exists
}

// Static verification that DefaultCommitTypeProvider implements CommitTypeProvider
var _ CommitTypeProvider = (*DefaultCommitTypeProvider)(nil)
