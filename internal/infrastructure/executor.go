package infrastructure

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/utils"
)

// CommandExecutor implémente l'interface domain.Executor
type CommandExecutor struct {
	log utils.Logger
}

// NewCommandExecutor crée une nouvelle instance de CommandExecutor
// Cette fonction servira de provider pour Wire
func NewCommandExecutor(logger utils.Logger) *CommandExecutor {
	return &CommandExecutor{
		log: logger,
	}
}

// Execute exécute une commande avec ses arguments et retourne la sortie
func (*CommandExecutor) Execute(ctx context.Context, cmd string, args ...string) (string, error) {
	command := exec.CommandContext(ctx, cmd, args...)
	output, err := command.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

// ExecuteWithEnv exécute une commande avec des variables d'environnement spécifiées
// et retourne la sortie
func (e *CommandExecutor) ExecuteWithEnv(
	ctx context.Context,
	env map[string]string,
	cmd string,
	args ...string,
) (string, error) {
	command := exec.CommandContext(ctx, cmd, args...)

	// Récupérer l'environnement actuel
	command.Env = os.Environ()

	// Ajouter les variables d'environnement spécifiées
	for key, value := range env {
		command.Env = append(command.Env, key+"="+value)
	}

	e.log.Debug("cmd", cmd)
	e.log.Debug("args", args)

	output, err := command.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

// Vérification statique que CommandExecutor implémente bien l'interface Executor
var _ utils.Executor = (*CommandExecutor)(nil)
