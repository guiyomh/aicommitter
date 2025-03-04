package ai

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/ollama/ollama/api"
)

// OllamaProvider is a implementation of AIProvider interface for Ollama
type OllamaProvider struct {
	baseURL     string
	model       string
	client      *api.Client
	maxDiffSize int
}

// NewOllamaProvider creates a new instance of OllamaProvider
func NewOllamaProvider(
	baseURL string,
	model string,
	maxDiffSize int,
) (*OllamaProvider, error) {
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	client := api.NewClient(parsedUrl, httpClient)
	return &OllamaProvider{
		baseURL:     baseURL,
		model:       model,
		client:      client,
		maxDiffSize: maxDiffSize,
	}, nil
}

func (p *OllamaProvider) GenerateCommitMessage(diff string, changedFiles []string) (string, error) {
	if len(diff) > p.maxDiffSize {
		diff = diff[:p.maxDiffSize] + "\n[diff truncated...]"
	}

	fileStr := strings.Join(changedFiles, ",")

	prompt := fmt.Sprintf(
		"Generates a commit message in conventional format for the following changes.\n\n"+
			"Files modified: %s\n\n"+
			"Diff:\n%s\n\n"+
			"Format: <type>(<scope>): <description>\n\n"+
			"Where <type> is one of the following: feat, fix, docs, style, refactor, perf, test, build, ci, chore, revert\n"+
			"<scope> is optional and represents the affected part of the code\n "+
			"<description> is a short description of the modifications\n\n",
		fileStr,
		diff,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var commitMessage string

	respFunc := func(resp api.GenerateResponse) error {
		commitMessage += resp.Response
		return nil
	}

	req := &api.GenerateRequest{
		Model:  p.model,
		Prompt: prompt,
		Stream: new(bool),
		// Temperature: 0.7,
		// TopK:        50,
		// TopP:        0.95,
	}

	err := p.client.Generate(ctx, req, respFunc)
	if err != nil {
		return "", fmt.Errorf("error while generating commit message with Ollama: %w", err)
	}

	return strings.TrimSpace(commitMessage), nil
}

// Vérification statique que OllamaProvider implémente bien l'interface AIProvider
var _ services.AIProvider = (*OllamaProvider)(nil)
