package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommitMessage_String(t *testing.T) {
	tests := []struct {
		name   string
		commit CommitMessage
		want   string
	}{
		{
			name: "message complet",
			commit: CommitMessage{
				Type:        "feat",
				Scope:       "auth",
				Description: "add google login",
				Body:        "Implement Google OAuth integration",
				Footer:      "Closes #123",
			},
			want: "feat(auth): add google login\n\nImplement Google OAuth integration\n\nCloses #123",
		},
		{
			name: "sans scope",
			commit: CommitMessage{
				Type:        "fix",
				Description: "correct typo",
				Body:        "Fix spelling mistake in README",
			},
			want: "fix: correct typo\n\nFix spelling mistake in README",
		},
		{
			name: "minimal",
			commit: CommitMessage{
				Type:        "chore",
				Description: "update dependencies",
			},
			want: "chore: update dependencies",
		},
		{
			name: "avec scope sans body",
			commit: CommitMessage{
				Type:        "feat",
				Scope:       "ui",
				Description: "add new button",
				Footer:      "BREAKING CHANGE",
			},
			want: "feat(ui): add new button\n\nBREAKING CHANGE",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.commit.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
