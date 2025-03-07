package services_test

import (
	"errors"
	"testing"

	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/guiyomh/aicommitter/internal/mocks/domain/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetDiff(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		expectedDiff := "diff --git a/file1 b/file1\nindex abc..def\n--- a/file1\n+++ b/file1"
		env := map[string]string{"GIT_PAGER": "cat"}
		mockExecutor.On("ExecuteWithEnv", mock.Anything, env, "git", "diff", "--diff-algorithm=minimal", ":(exclude)*.lock", ":(exclude)*.sum").Return(expectedDiff, nil)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		diff, err := diffService.GetDiff()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedDiff, diff)
		mockExecutor.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		expectedErr := errors.New("command failed")
		env := map[string]string{"GIT_PAGER": "cat"}
		mockExecutor.On("ExecuteWithEnv", mock.Anything, env, "git", "diff", "--diff-algorithm=minimal", ":(exclude)*.lock", ":(exclude)*.sum").Return("", expectedErr)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		diff, err := diffService.GetDiff()

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get git diff")
		assert.Empty(t, diff)
		mockExecutor.AssertExpectations(t)
	})
}

func TestGetStagedDiff(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		expectedDiff := "diff --git a/file1 b/file1\nindex abc..def\n--- a/file1\n+++ b/file1"
		env := map[string]string{"GIT_PAGER": "cat"}
		mockExecutor.On("ExecuteWithEnv", mock.Anything, env, "git", "diff", "--staged", "--diff-algorithm=minimal", ":(exclude)*.lock", ":(exclude)*.sum").Return(expectedDiff, nil)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		diff, err := diffService.GetStagedDiff()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedDiff, diff)
		mockExecutor.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		expectedErr := errors.New("command failed")
		env := map[string]string{"GIT_PAGER": "cat"}
		mockExecutor.On("ExecuteWithEnv", mock.Anything, env, "git", "diff", "--staged", "--diff-algorithm=minimal", ":(exclude)*.lock", ":(exclude)*.sum").Return("", expectedErr)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		diff, err := diffService.GetStagedDiff()

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get staged git diff")
		assert.Empty(t, diff)
		mockExecutor.AssertExpectations(t)
	})
}

func TestGetFilesChanged(t *testing.T) {
	t.Run("success with files", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		expectedOutput := "file1.go\nfile2.go\nfile3.go"
		expectedFiles := []string{"file1.go", "file2.go", "file3.go"}
		mockExecutor.On("Execute", mock.Anything, "git", "diff", "--name-only").Return(expectedOutput, nil)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		files, err := diffService.GetFilesChanged()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedFiles, files)
		mockExecutor.AssertExpectations(t)
	})

	t.Run("success with no files", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		mockExecutor.On("Execute", mock.Anything, "git", "diff", "--name-only").Return("", nil)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		files, err := diffService.GetFilesChanged()

		// Assert
		assert.NoError(t, err)
		assert.Empty(t, files)
		mockExecutor.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		expectedErr := errors.New("command failed")
		mockExecutor.On("Execute", mock.Anything, "git", "diff", "--name-only").Return("", expectedErr)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		files, err := diffService.GetFilesChanged()

		// Assert
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to get changed files")
		assert.Nil(t, files)
		mockExecutor.AssertExpectations(t)
	})

	t.Run("success with single file", func(t *testing.T) {
		// Arrange
		mockExecutor := new(utils.MockExecutor)
		expectedOutput := "file1.go"
		expectedFiles := []string{"file1.go"}
		mockExecutor.On("Execute", mock.Anything, "git", "diff", "--name-only").Return(expectedOutput, nil)

		diffService := services.NewDiffService(mockExecutor)

		// Act
		files, err := diffService.GetFilesChanged()

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedFiles, files)
		mockExecutor.AssertExpectations(t)
	})
}
