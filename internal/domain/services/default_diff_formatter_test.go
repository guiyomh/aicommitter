package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultDiffFormatter_FormatDiff(t *testing.T) {
	formatter := NewDefaultDiffFormatter()

	tests := []struct {
		name     string
		diff     string
		maxSize  int
		expected string
	}{
		{
			name:     "diff plus court que maxSize",
			diff:     "modification simple",
			maxSize:  20,
			expected: "modification simple",
		},
		{
			name:     "diff plus long que maxSize",
			diff:     "ceci est une très longue modification qui devrait être tronquée",
			maxSize:  20,
			expected: "ceci est une très lo\n[diff truncated...]",
		},
		{
			name:     "maxSize égal à la longueur du diff",
			diff:     "test exact",
			maxSize:  9,
			expected: "test exact",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatter.FormatDiff(tt.diff, tt.maxSize)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestDefaultDiffFormatter_GetFileList(t *testing.T) {
	formatter := NewDefaultDiffFormatter()

	tests := []struct {
		name     string
		files    []string
		expected string
	}{
		{
			name:     "liste vide",
			files:    []string{},
			expected: "",
		},
		{
			name:     "un seul fichier",
			files:    []string{"fichier1.go"},
			expected: "fichier1.go",
		},
		{
			name:     "plusieurs fichiers",
			files:    []string{"fichier1.go", "fichier2.go", "fichier3.go"},
			expected: "fichier1.go, fichier2.go, fichier3.go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatter.GetFileList(tt.files)
			assert.Equal(t, tt.expected, result)
		})
	}
}
