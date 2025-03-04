package config

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
		MaxDiffSize        int                   `mapstructure:"max_diff_size"`
	} `mapstructure:"git"`
	AI struct {
		Provider string `mapstructure:"provider"`
		Model    string `mapstructure:"model"`
		BaseURL  string `mapstructure:"base_url"`
		APIKey   string `mapstructure:"api_key"`
	} `mapstructure:"ai"`
	Log struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"log"`
}
