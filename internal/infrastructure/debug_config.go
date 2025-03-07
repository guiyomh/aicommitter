package infrastructure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// DebugConfigPaths affiche tous les chemins de configuration recherchés
func DebugConfigPaths() {
	fmt.Println("Debugging config paths:")

	// Créer une instance viper avec les mêmes paramètres
	v := viper.New()
	v.SetConfigName("aicommitter")
	v.SetConfigType("yaml")

	paths := []string{
		".",
		os.ExpandEnv("$HOME/.config/aicommitter"),
		"/etc/aicommitter",
	}

	for _, path := range paths {
		expandedPath := os.ExpandEnv(path)
		fullPath := filepath.Join(expandedPath, "aicommitter.yaml")

		// Vérifier si le fichier existe
		_, err := os.Stat(fullPath)
		status := "NOT FOUND"
		if err == nil {
			status = "EXISTS"
		}

		fmt.Printf("- %s: %s\n", fullPath, status)
	}

	// Vérifier les variables d'environnement
	fmt.Println("\nEnvironment variables:")
	for _, env := range os.Environ() {
		if strings.HasPrefix(env, "AICOMMITTER_") {
			fmt.Println("-", env)
		}
	}
}
