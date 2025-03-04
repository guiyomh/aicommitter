package infrastructure

import (
	"context"
	"os/exec"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/utils"
)

// CommandExecutor implémente l'interface domain.Executor
type CommandExecutor struct{}

// NewCommandExecutor crée une nouvelle instance de CommandExecutor
// Cette fonction servira de provider pour Wire
func NewCommandExecutor() *CommandExecutor {
	return &CommandExecutor{}
}

// Execute exécute une commande avec ses arguments et retourne la sortie
func (e *CommandExecutor) Execute(ctx context.Context, cmd string, args ...string) (string, error) {
	command := exec.CommandContext(ctx, cmd, args...)
	output, err := command.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

// Vérification statique que CommandExecutor implémente bien l'interface Executor
var _ utils.Executor = (*CommandExecutor)(nil)
