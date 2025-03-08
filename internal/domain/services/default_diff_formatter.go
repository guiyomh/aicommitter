package services

import "strings"

// DefaultDiffFormatter implémente DiffFormatter
type DefaultDiffFormatter struct{}

func NewDefaultDiffFormatter() *DefaultDiffFormatter {
	return &DefaultDiffFormatter{}
}

func (*DefaultDiffFormatter) FormatDiff(diff string, maxSize int) string {
	if len(diff) > maxSize+1 {
		return diff[:maxSize+1] + "\n[diff truncated...]"
	}
	return diff
}

func (*DefaultDiffFormatter) GetFileList(files []string) string {
	return strings.Join(files, ", ")
}

// Static verification that DefaultDiffFormatter implements DiffFormatter
var _ DiffFormatter = (*DefaultDiffFormatter)(nil)
