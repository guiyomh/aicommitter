// Package conventionalcommit provides a parser for Conventional Commits specification.
// For more information about the specification: https://www.conventionalcommits.org
package conventionalcommit

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Parser définit l'interface pour analyser les messages de commit conventionnels
type Parser interface {
	Parse(message string) (*Commit, error)
}

// DefaultParser est l'implémentation par défaut du Parser
type DefaultParser struct{}

// NewParser crée une nouvelle instance de DefaultParser
func NewParser() Parser {
	return &DefaultParser{}
}

// ValidationError represents an error that occurs during commit message validation
type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

// IsValidationError checks if an error is a ValidationError
func IsValidationError(err error) bool {
	var validationErr *ValidationError
	return errors.As(err, &validationErr)
}

var (
	// commitPattern defines the standard format of a conventional commit
	// Format: <type>[optional scope]: <description>
	// [optional body]
	// [optional footer(s)]
	commitPattern = regexp.MustCompile(`^(?P<type>\w+)(?:\((?P<scope>[\w-]+)\))?!?: (?P<description>.+)`)

	// ErrInvalidFormat is returned when the commit message format is invalid
	ErrInvalidFormat = &ValidationError{Message: "invalid conventional commit format"}
)

// Type represents the conventional commit type
type Type string

// Standard commit types
const (
	TypeFeat     Type = "feat"     // New features
	TypeFix      Type = "fix"      // Bug fixes
	TypeDocs     Type = "docs"     // Documentation changes
	TypeStyle    Type = "style"    // Code style changes (formatting, spacing, etc.)
	TypeRefactor Type = "refactor" // Code refactoring
	TypePerf     Type = "perf"     // Performance improvements
	TypeTest     Type = "test"     // Adding or modifying tests
	TypeChore    Type = "chore"    // Maintenance tasks
)

// IsValid checks if the commit type is valid
func (t Type) IsValid() bool {
	switch t {
	case TypeFeat, TypeFix, TypeDocs, TypeStyle, TypeRefactor, TypePerf, TypeTest, TypeChore:
		return true
	default:
		return false
	}
}

// Commit represents a parsed conventional commit
type Commit struct {
	Type        Type   `json:"type"`        // Commit type (feat, fix, etc.)
	Scope       string `json:"scope"`       // Change scope (optional)
	Description string `json:"description"` // Short description of the change
	Body        string `json:"body"`        // Message body (optional)
	Footer      string `json:"footer"`      // Message footer (optional)
	Breaking    bool   `json:"breaking"`    // Indicates if this is a breaking change
}

// Parse analyzes a commit message and returns a Commit structure.
// Returns an error if the message format is invalid.
func (p *DefaultParser) Parse(message string) (*Commit, error) {
	lines := strings.Split(message, "\n")
	if len(lines) == 0 {
		return nil, fmt.Errorf("empty commit message: %w", ErrInvalidFormat)
	}

	// Parse the first line (header)
	matches := commitPattern.FindStringSubmatch(lines[0])
	if matches == nil {
		return nil, fmt.Errorf("malformed commit message header: %w", ErrInvalidFormat)
	}

	commitType := Type(strings.ToLower(matches[1]))
	if !commitType.IsValid() {
		return nil, &ValidationError{Message: fmt.Sprintf("invalid commit type: %s", commitType)}
	}

	commit := &Commit{
		Type:        commitType,
		Scope:       matches[2],
		Description: matches[3],
		Breaking:    strings.Contains(lines[0], "!"),
	}

	// Parse body and footer if present
	if len(lines) > 1 {
		bodyStart := 1
		for ; bodyStart < len(lines) && len(strings.TrimSpace(lines[bodyStart])) == 0; bodyStart++ {
		}

		if bodyStart < len(lines) {
			// Look for footer (starts with "BREAKING CHANGE:" or other keyword followed by ":")
			footerStart := -1
			for i := bodyStart; i < len(lines); i++ {
				line := strings.TrimSpace(lines[i])
				if strings.HasPrefix(line, "BREAKING CHANGE:") ||
					(strings.Contains(line, ":") && !strings.Contains(line, " :")) {
					footerStart = i
					break
				}
			}

			if footerStart != -1 {
				commit.Body = strings.TrimSpace(strings.Join(lines[bodyStart:footerStart], "\n"))
				commit.Footer = strings.TrimSpace(strings.Join(lines[footerStart:], "\n"))
				if strings.Contains(commit.Footer, "BREAKING CHANGE:") {
					commit.Breaking = true
				}
			} else {
				commit.Body = strings.TrimSpace(strings.Join(lines[bodyStart:], "\n"))
			}
		}
	}

	return commit, nil
}

// String returns a string representation of the commit following the conventional commit format
func (c *Commit) String() string {
	var builder strings.Builder

	// Header
	builder.WriteString(string(c.Type))
	if c.Scope != "" {
		builder.WriteString("(" + c.Scope + ")")
	}
	if c.Breaking {
		builder.WriteString("!")
	}
	builder.WriteString(": " + c.Description)

	// Body
	if c.Body != "" {
		builder.WriteString("\n\n" + c.Body)
	}

	// Footer
	if c.Footer != "" {
		builder.WriteString("\n\n" + c.Footer)
	}

	return builder.String()
}
