package infrastructure

import (
	"fmt"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/spf13/viper"
)

func LoadConfig() (*domain.Config, error) {
	v := viper.New()

	setDefaultValues(v)

	// Configuration des sources et format
	v.SetConfigName("aicommiter")
	v.SetConfigType("yaml")

	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.config/aicommitter")
	v.AddConfigPath("/etc/aicommitter")

	v.SetEnvPrefix("AICOMMITTER")
	v.AutomaticEnv()

	_ = v.ReadInConfig()

	var cfg domain.Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error when reading configuration: %w", err)
	}

	return &cfg, nil
}

// Ajouter cette fonction à la fin du fichier

// WriteDefaultConfig crée un fichier de configuration avec les valeurs par défaut
func WriteDefaultConfig(path string) error {
	v := viper.New()

	// Définir les valeurs par défaut (mêmes que dans LoadConfig)
	setDefaultValues(v)

	// Configurer Viper pour le fichier de sortie
	v.SetConfigName("aicommiter")
	v.SetConfigType("yaml")

	// Écrire la configuration dans un fichier
	err := v.WriteConfigAs(path)
	if err != nil {
		return fmt.Errorf("impossible d'écrire la configuration dans %s: %w", path, err)
	}

	return nil
}

// setDefaultValues définit les valeurs par défaut pour la configuration
func setDefaultValues(v *viper.Viper) {
	v.SetDefault("git.default_commit_style", "conventional")
	v.SetDefault("git.commit_types", map[string]interface{}{
		"feat": map[string]string{
			"title":       "Features",
			"description": "A new feature",
			"emoji":       "✨",
		},
		"fix": map[string]string{
			"title":       "Bug Fixes",
			"description": "A bug Fix",
			"emoji":       "🐛",
		},
		"docs": map[string]string{
			"title":       "Documentation",
			"description": "Documentation only changes",
			"emoji":       "📚",
		},
		"style": map[string]string{
			"title":       "Styles",
			"description": "Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)",
			"emoji":       "💎",
		},
		"refactor": map[string]string{
			"title":       "Code Refactoring",
			"description": "A code change that neither fixes a bug nor adds a feature",
			"emoji":       "📦",
		},
		"perf": map[string]string{
			"title":       "Performance Improvements",
			"description": "A code change that improves performance",
			"emoji":       "🚀",
		},
		"test": map[string]string{
			"title":       "Tests",
			"description": "Adding missing tests or correcting existing tests",
			"emoji":       "🚨",
		},
		"build": map[string]string{
			"title":       "Builds",
			"description": "Changes that affect the build system or external dependencies",
			"emoji":       "🛠",
		},
		"ci": map[string]string{
			"title":       "Continuous Integrations",
			"description": "Changes to CI configuration files and scripts",
			"emoji":       "⚙️",
		},
		"chore": map[string]string{
			"title":       "Chores",
			"description": "Other changes that don't modify src or test files",
			"emoji":       "♻️",
		},
		"revert": map[string]string{
			"title":       "Reverts",
			"description": "Reverts a previous commit",
			"emoji":       "🗑",
		},
	})
	v.SetDefault("log.level", "info")
}
