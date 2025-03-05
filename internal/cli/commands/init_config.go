package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/guiyomh/aicommitter/internal/infrastructure"
	"github.com/spf13/cobra"
)

func InitConfigCommand() *cobra.Command {
	var force bool
	var globalConfig bool

	init := &cobra.Command{
		Use:   "init",
		Short: "Initializes application configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			var configPath string
			if globalConfig {
				homedir, err := os.UserHomeDir()
				if err != nil {
					return fmt.Errorf("could not get home directory: %w", err)
				}
				configPath = filepath.Join(homedir, ".config", "aicommitter")
				if err := os.MkdirAll(configPath, 0755); err != nil {
					return fmt.Errorf("could not create configuration directory %s: %w", configPath, err)
				}
				configPath = filepath.Join(configPath, ".aicommitter.yaml")

			} else {
				configPath = ".aicommitter.yaml"
			}

			if _, err := os.Stat(configPath); err == nil && !force {
				return fmt.Errorf("configuration file %s already exists, use --force to overwrite", configPath)
			}

			if err := infrastructure.WriteDefaultConfig(configPath); err != nil {
				return fmt.Errorf("could not write configuration file: %w", err)
			}

			fmt.Printf("✅ Configuration file written to %s\n", configPath)
			return nil
		},
	}

	init.Flags().BoolVarP(&force, "force", "f", false, "Overwrite existing configuration file")
	init.Flags().BoolVarP(&globalConfig, "global", "g", false, "Create configuration in $HOME/.config/ rather than in the current directory")

	return init
}
