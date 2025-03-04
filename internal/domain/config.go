package domain

// CommitType represents a commit type with its emoji and description
type CommitType struct {
	Title       string `mapstructure:"title"`
	Description string `mapstructure:"description"`
	Emoji       string `mapstructure:"emoji"`
}

// Config represents the configuration of the application
type Config struct {
	Git struct {
		DefaultCommitStyle string                `mapstructure:"default_commit_style"`
		CommitTypes        map[string]CommitType `mapstructure:"commit_types"`
	} `mapstructure:"git"`
	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
}
