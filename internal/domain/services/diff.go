package services

import (
	"context"
	"errors"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/utils"
)

// DiffService provides functionality for working with git diffs
type DiffService struct {
	executor utils.Executor
}

// NewDiffService creates a new DiffService instance
func NewDiffService(executor utils.Executor) *DiffService {
	return &DiffService{
		executor: executor,
	}
}

// executeDiffCommand exécute la commande git diff avec les options spécifiées
func (s *DiffService) executeDiffCommand(staged bool, args ...string) (string, error) {
	env := map[string]string{
		"GIT_PAGER": "cat",
	}

	cmdArgs := []string{"diff"}
	if staged {
		cmdArgs = append(cmdArgs, "--staged")
	}
	cmdArgs = append(cmdArgs, "--diff-algorithm=minimal", ":(exclude)*.lock", ":(exclude)*.sum")
	cmdArgs = append(cmdArgs, args...)

	output, err := s.executor.ExecuteWithEnv(context.Background(), env, "git", cmdArgs...)
	if err != nil {
		action := "get git diff"
		if staged {
			action = "get staged git diff"
		}
		return "", errors.New("failed to " + action + ": " + err.Error())
	}
	return output, nil
}

// GetDiff returns the git diff for the current changes
func (s *DiffService) GetDiff() (string, error) {
	return s.executeDiffCommand(false)
}

// GetStagedDiff returns the git diff for staged changes
func (s *DiffService) GetStagedDiff() (string, error) {
	return s.executeDiffCommand(true)
}

// GetFilesChanged returns a list of files that have been changed
func (s *DiffService) GetFilesChanged() ([]string, error) {
	output, err := s.executor.Execute(context.Background(), "git", "diff", "--name-only")
	if err != nil {
		return nil, errors.New("failed to get changed files: " + err.Error())
	}

	files := strings.Split(strings.TrimSpace(output), "\n")
	if len(files) == 1 && files[0] == "" {
		return []string{}, nil
	}
	return files, nil
}
