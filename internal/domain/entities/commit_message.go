package entities

// CommitMessage représente un message de commit conventionnel
type CommitMessage struct {
	Type        string `json:"type"`
	Scope       string `json:"scope,omitempty"`
	Description string `json:"description"`
	Body        string `json:"body,omitempty"`
	Footer      string `json:"footer,omitempty"`
	Breaking    bool   `json:"breaking,omitempty"`
}

// String retourne une représentation en chaîne du message de commit
func (c *CommitMessage) String() string {
	var result string

	// Type et scope
	result = c.Type
	if c.Scope != "" {
		result += "(" + c.Scope + ")"
	}

	// Breaking change
	if c.Breaking {
		result += "!"
	}

	// Description
	result += ": " + c.Description

	// Body
	if c.Body != "" {
		result += "\n\n" + c.Body
	}

	// Footer
	if c.Footer != "" {
		result += "\n\n" + c.Footer
	}

	return result
}
