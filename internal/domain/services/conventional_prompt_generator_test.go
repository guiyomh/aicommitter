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
	mockDiffFormatter := new(services.MockDiffFormatter)

	// Act
	generator := NewConventionalPromptGenerator(mockTypeProvider, mockTemplateGen, mockDiffFormatter)

	// Assert
	assert.NotNil(t, generator)
	assert.Equal(t, mockTypeProvider, generator.typeProvider)
	assert.Equal(t, mockTemplateGen, generator.templateGen)
	assert.Equal(t, mockDiffFormatter, generator.diffFormatter)
}

func TestGenerateCommitPrompt(t *testing.T) {
	// Arrange
	mockTypeProvider := new(services.MockCommitTypeProvider)
	mockTemplateGen := new(services.MockPromptTemplateGenerator)
	mockDiffFormatter := new(services.MockDiffFormatter)

	generator := NewConventionalPromptGenerator(mockTypeProvider, mockTemplateGen, mockDiffFormatter)

	testCases := []struct {
		name          string
		diff          string
		changedFiles  []string
		maxDiffSize   int
		commitTypes   map[string]config.CommitType
		template      string
		formattedDiff string
		fileList      string
		expected      string
	}{
		{
			name:         "Basic commit prompt generation",
			diff:         "test diff",
			changedFiles: []string{"file1.go", "file2.go"},
			maxDiffSize:  1000,
			commitTypes: map[string]config.CommitType{
				"feat": {Description: "Features"},
				"fix":  {Description: "Bug Fixes"},
			},
			template:      "Template %s %s",
			formattedDiff: "formatted diff",
			fileList:      "file1.go, file2.go",
			expected:      "Template formatted diff file1.go, file2.go",
		},
		{
			name:         "Empty diff",
			diff:         "",
			changedFiles: []string{"file1.go"},
			maxDiffSize:  1000,
			commitTypes: map[string]config.CommitType{
				"feat": {Description: "Features"},
			},
			template:      "Template %s %s",
			formattedDiff: "",
			fileList:      "file1.go",
			expected:      "Template  file1.go",
		},
		{
			name:         "Multiple commit types",
			diff:         "test diff",
			changedFiles: []string{"file1.go"},
			maxDiffSize:  500,
			commitTypes: map[string]config.CommitType{
				"feat":     {Description: "Features"},
				"fix":      {Description: "Bug Fixes"},
				"docs":     {Description: "Documentation"},
				"refactor": {Description: "Code Refactoring"},
			},
			template:      "Complex Template %s %s",
			formattedDiff: "truncated diff",
			fileList:      "file1.go",
			expected:      "Complex Template truncated diff file1.go",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Configure mocks
			mockTypeProvider.On("GetCommitTypes").Return(tc.commitTypes).Once()
			mockDiffFormatter.On("FormatDiff", tc.diff, tc.maxDiffSize).Return(tc.formattedDiff).Once()
			mockDiffFormatter.On("GetFileList", tc.changedFiles).Return(tc.fileList).Once()
			mockTemplateGen.On("GenerateTemplate", tc.commitTypes).Return(tc.template).Once()

			// Act
			result := generator.GenerateCommitPrompt(tc.diff, tc.changedFiles, tc.maxDiffSize)

			// Assert
			assert.Equal(t, tc.expected, result)

			// Verify mock calls
			mockTypeProvider.AssertExpectations(t)
			mockTemplateGen.AssertExpectations(t)
			mockDiffFormatter.AssertExpectations(t)
		})
	}
}
