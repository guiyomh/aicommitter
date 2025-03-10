package models

import "strings"

// CommitMessage represents a structured commit message following the Conventional Commits specification
type CommitMessage struct {
	Type        string
	Scope       string
	Description string
	Body        string
	Footer      string
	Breaking    bool
}

// String formats the commit message according to conventional commits specification
func (m *CommitMessage) String() string {
	var parts []string

	// Build the header
	header := m.Type
	if m.Scope != "" {
		header += "(" + m.Scope + ")"
	}
	if m.Breaking {
		header += "!"
	}
	header += ": " + m.Description
	parts = append(parts, header)

	// Add body if present
	if m.Body != "" {
		parts = append(parts, "", m.Body)
	}

	// Add footer if present
	if m.Footer != "" {
		parts = append(parts, "", m.Footer)
	}

	return strings.Join(parts, "\n")
}
