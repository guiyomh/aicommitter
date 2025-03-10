package services

import (
	"testing"

	"github.com/guiyomh/aicommitter/internal/domain/config"
	"github.com/guiyomh/aicommitter/internal/mocks/services"

	"github.com/stretchr/testify/assert"
)

func TestNewConventionalPromptGenerator(t *testing.T) {
	// Arrange
	mockTypeProvider := new(services.MockCommitTypeProvider)
	mockTemplateGen := new(services.MockPromptTemplateGenerator)

	// Act
	generator := NewConventionalPromptGenerator(mockTypeProvider, mockTemplateGen)

	// Assert
	assert.NotNil(t, generator)
	assert.Equal(t, mockTypeProvider, generator.typeProvider)
	assert.Equal(t, mockTemplateGen, generator.templateGen)
}

func TestGenerateCommitPrompt(t *testing.T) {
	// Arrange
	mockTypeProvider := new(services.MockCommitTypeProvider)
	mockTemplateGen := new(services.MockPromptTemplateGenerator)

	generator := NewConventionalPromptGenerator(mockTypeProvider, mockTemplateGen)

	testCases := []struct {
		name        string
		maxDiffSize int
		commitTypes map[string]config.CommitType
		template    string
		expected    string
	}{
		{
			name:        "Basic commit prompt generation",
			maxDiffSize: 1000,
			commitTypes: map[string]config.CommitType{
				"feat": {Description: "Features"},
				"fix":  {Description: "Bug Fixes"},
			},
			template: "Template",
			expected: "Template",
		},
		{
			name:        "Empty commit types",
			maxDiffSize: 1000,
			commitTypes: map[string]config.CommitType{},
			template:    "Template",
			expected:    "Template",
		},
		{
			name:        "Multiple commit types",
			maxDiffSize: 500,
			commitTypes: map[string]config.CommitType{
				"feat":     {Description: "Features"},
				"fix":      {Description: "Bug Fixes"},
				"docs":     {Description: "Documentation"},
				"refactor": {Description: "Code Refactoring"},
			},
			template: "Complex Template",
			expected: "Complex Template",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Configure mocks
			mockTypeProvider.On("GetCommitTypes").Return(tc.commitTypes).Once()
			mockTemplateGen.On("GenerateTemplate", tc.commitTypes).Return(tc.template).Once()

			// Act
			result := generator.GenerateCommitPrompt(tc.maxDiffSize)

			// Assert
			assert.Equal(t, tc.expected, result)

			// Verify mock calls
			mockTypeProvider.AssertExpectations(t)
			mockTemplateGen.AssertExpectations(t)
		})
	}
}
