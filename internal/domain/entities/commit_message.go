package entities

import "fmt"

// CommitMessage représente la structure d'un message de commit conventionnel
type CommitMessage struct {
	Type        string `json:"type"`
	Scope       string `json:"scope,omitempty"`
	Description string `json:"description"`
	Body        string `json:"body,omitempty"`
	Footer      string `json:"footer,omitempty"`
}

// String retourne le message de commit formaté selon les conventions
func (c *CommitMessage) String() string {
	var commitMsg string

	// Construction de la première ligne (type(scope): description)
	if c.Scope != "" {
		commitMsg = fmt.Sprintf("%s(%s): %s", c.Type, c.Scope, c.Description)
	} else {
		commitMsg = fmt.Sprintf("%s: %s", c.Type, c.Description)
	}

	// Ajout du body s'il existe
	if c.Body != "" {
		commitMsg += "\n\n" + c.Body
	}

	// Ajout du footer s'il existe
	if c.Footer != "" {
		commitMsg += "\n\n" + c.Footer
	}

	return commitMsg
}
