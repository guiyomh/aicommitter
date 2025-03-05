package services

import (
	"strings"

	"github.com/guiyomh/aicommitter/internal/domain/config"
)

// ConventionalPromptGenerator implémente l'interface PromptGenerator avec une logique par défaut
type ConventionalPromptGenerator struct {
	commitTypes map[string]config.CommitType
}

// NewDefaultPromptGenerator crée une nouvelle instance de DefaultPromptGenerator
func NewDefaultPromptGenerator(cfg *config.Config) *ConventionalPromptGenerator {
	return &ConventionalPromptGenerator{
		commitTypes: cfg.Git.CommitTypes,
	}
}

// GenerateCommitPrompt génère un prompt pour la création de messages de commit
func (g *ConventionalPromptGenerator) GenerateCommitPrompt(diff string, changedFiles []string, maxDiffSize int) string {
	if len(diff) > maxDiffSize {
		diff = diff[:maxDiffSize] + "\n[diff truncated...]"
	}

	fileStr := strings.Join(changedFiles, ",")

	return `Tu es un assistant spécialisé dans la génération de messages de commit au format Conventional Commits. 

# Objectif
Générer un message de commit descriptif, clair et concis qui suit strictement la spécification Conventional Commits.

# Contexte du Commit

- Diff Git: 

` + "```\n" + diff + "\n```" + `

- Fichiers impactés: ` + fileStr + `

# Contraintes et Instructions

## Format du Commit

Le message DOIT suivre le format :

` + "```" + `
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
` + "```" + `

## Règles de Génération

1. Choix du Type de Commit

- Utiliser l'un des types standard :

  - feat: nouvelle fonctionnalité
  - fix: correction de bug
  - docs: modifications de documentation
  - style: formatage, point-virgules manquants, etc.
  - refactor: refactorisation du code
  - test: ajout/modification de tests
  - chore: maintenance, mises à jour de dépendances
  - perf: amélioration des performances
  - ci: modifications des configurations CI

2. Scope (Optionnel)

- Indiquer le composant/module/fichier principalement affecté
- Sera limité à une portée technique ou fonctionnelle

3. Description

- Courte (< 50 caractères)
- Impérative, style "Ajouter/Corriger/Modifier..."
- Commencer par une lettre minuscule
- Pas de point final

4. Corps du Message (Optionnel)

- Explication plus détaillée si nécessaire
- Motivation du changement
- Différences avec l'état précédent

5. Pieds de Page (Optionnel)

- Références d'issues (ex: "Fixes #123")
- Méta-informations
` +

		// # Paramètres Spécifiques Fournis

		// - Type forcé : {type_force}
		// - Scope forcé : {scope_force}
		// - Numéro d'issue forcé : {issue_force}
		// - Langue du message : {langue}

		`
# Consignes Supplémentaires

- Soyez concis et précis
- Utilisez le contexte du diff pour comprendre les changements
- Si aucun type n'est évident, choisissez le plus approprié
- En cas de doute, préférez refactor ou chore

# Format de Réponse
Répondez UNIQUEMENT avec le message de commit, sans aucun texte supplémentaire.
Le message DOIT être facilement parsable, donc :

- Utilisez un délimiteur clair comme ---COMMIT_MESSAGE_START--- et ---COMMIT_MESSAGE_END---
- Incluez des métadonnées parsables

Exemple de format de réponse :

` + "```" + `
---COMMIT_MESSAGE_START---
{
  "type": "feat",
  "scope": "authentication",
  "description": "add login with google oauth",
  "body": "Implement Google OAuth integration for user authentication\nAdds support for Google sign-in on the login page",
  "footer": "Fixes #456"
}
---COMMIT_MESSAGE_END---
` + "```\n"
}

// Vérification statique que DefaultPromptGenerator implémente bien l'interface PromptGenerator
var _ PromptGenerator = (*ConventionalPromptGenerator)(nil)
