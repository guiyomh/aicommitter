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

// GetDiff returns the git diff for the current changes
func (s *DiffService) GetDiff() (string, error) {
	env := map[string]string{
		"GIT_PAGER": "cat",
	}
	output, err := s.executor.ExecuteWithEnv(context.Background(),
		env,
		"git",
		"diff",
		"--diff-algorithm=minimal",
		":(exclude)*.lock",
		":(exclude)*.sum",
	)
	if err != nil {
		return "", errors.New("failed to get git diff: " + err.Error())
	}
	return output, nil
}

// GetStagedDiff returns the git diff for staged changes
func (s *DiffService) GetStagedDiff() (string, error) {
	env := map[string]string{
		"GIT_PAGER": "cat",
	}
	output, err := s.executor.ExecuteWithEnv(context.Background(),
		env,
		"git",
		"diff",
		"--staged",
		"--diff-algorithm=minimal",
		":(exclude)*.lock",
		":(exclude)*.sum",
	)
	if err != nil {
		return "", errors.New("failed to get staged git diff: " + err.Error())
	}
	return output, nil
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
