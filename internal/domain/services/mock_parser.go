package services

import (
	"github.com/guiyomh/aicommitter/pkg/conventionalcommit"
)

// MockParser est une implémentation mock du Parser pour les tests
type MockParser struct {
	ParseFunc func(message string) (*conventionalcommit.Commit, error)
}

func (m *MockParser) Parse(message string) (*conventionalcommit.Commit, error) {
	if m.ParseFunc != nil {
		return m.ParseFunc(message)
	}
	return nil, nil
}
