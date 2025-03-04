package commands

import (
	"fmt"

	"github.com/guiyomh/aicommitter/internal/domain"
	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/spf13/cobra"
)

func CommitCommand(diffService *services.DiffService, cfg *domain.Config) *cobra.Command {
	commit := &cobra.Command{
		Use:   "commit",
		Short: "Commit a model",
		RunE: func(_ *cobra.Command, _ []string) error {
			diff, err := diffService.GetStagedDiff()
			if err != nil {
				return fmt.Errorf("erreur lors de la récupération des modifications : %w", err)
			}

			// Afficher le diff (vous pouvez faire d'autres opérations ici)
			fmt.Println("Modifications à commiter :")
			fmt.Println(diff)

			fmt.Println("Niveau de log :")
			fmt.Println(cfg.Log.Level)

			return nil
		},
	}
	return commit
}
