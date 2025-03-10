package conventionalcommit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		message string
		want    *Commit
		wantErr bool
	}{
		{
			name:    "invalid format",
			message: "invalid message",
			wantErr: true,
		},
		{
			name:    "empty message",
			message: "",
			wantErr: true,
		},
		{
			name:    "invalid type",
			message: "unknown: some message",
			wantErr: true,
		},
		{
			name:    "simple commit",
			message: "feat: add new feature",
			want: &Commit{
				Type:        TypeFeat,
				Description: "add new feature",
			},
		},
		{
			name:    "commit with scope",
			message: "fix(auth): fix login issue",
			want: &Commit{
				Type:        TypeFix,
				Scope:       "auth",
				Description: "fix login issue",
			},
		},
		{
			name:    "commit with breaking change",
			message: "feat!: change API\n\nBREAKING CHANGE: new API version",
			want: &Commit{
				Type:        TypeFeat,
				Description: "change API",
				Footer:      "BREAKING CHANGE: new API version",
				Breaking:    true,
			},
		},
		{
			name: "complete commit",
			message: `feat(api)!: major API change

Detailed description of the changes
Over multiple lines

BREAKING CHANGE: incompatible change
Signed-off-by: John Doe <john@example.com>`,
			want: &Commit{
				Type:        TypeFeat,
				Scope:       "api",
				Description: "major API change",
				Body:        "Detailed description of the changes\nOver multiple lines",
				Footer:      "BREAKING CHANGE: incompatible change\nSigned-off-by: John Doe <john@example.com>",
				Breaking:    true,
			},
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.Parse(tt.message)
			if tt.wantErr {
				require.Error(t, err)
				assert.True(t, IsValidationError(err), "error should be a ValidationError")
				return
			}

			require.NoError(t, err)
			require.NotNil(t, got)
			assert.Equal(t, tt.want.Type, got.Type)
			assert.Equal(t, tt.want.Scope, got.Scope)
			assert.Equal(t, tt.want.Description, got.Description)
			assert.Equal(t, tt.want.Body, got.Body)
			assert.Equal(t, tt.want.Footer, got.Footer)
			assert.Equal(t, tt.want.Breaking, got.Breaking)
		})
	}
}

func TestType_IsValid(t *testing.T) {
	tests := []struct {
		name string
		t    Type
		want bool
	}{
		{"feat is valid", TypeFeat, true},
		{"fix is valid", TypeFix, true},
		{"docs is valid", TypeDocs, true},
		{"style is valid", TypeStyle, true},
		{"refactor is valid", TypeRefactor, true},
		{"perf is valid", TypePerf, true},
		{"test is valid", TypeTest, true},
		{"chore is valid", TypeChore, true},
		{"unknown is invalid", "unknown", false},
		{"empty is invalid", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.t.IsValid())
		})
	}
}

func TestCommit_String(t *testing.T) {
	tests := []struct {
		name   string
		commit *Commit
		want   string
	}{
		{
			name: "simple commit",
			commit: &Commit{
				Type:        TypeFeat,
				Description: "add new feature",
			},
			want: "feat: add new feature",
		},
		{
			name: "commit with scope",
			commit: &Commit{
				Type:        TypeFix,
				Scope:       "auth",
				Description: "fix login issue",
			},
			want: "fix(auth): fix login issue",
		},
		{
			name: "commit with breaking change",
			commit: &Commit{
				Type:        TypeFeat,
				Description: "change API",
				Breaking:    true,
				Footer:      "BREAKING CHANGE: new API version",
			},
			want: "feat!: change API\n\nBREAKING CHANGE: new API version",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.commit.String())
		})
	}
}
