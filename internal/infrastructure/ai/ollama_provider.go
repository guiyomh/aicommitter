package ai

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
	"github.com/ollama/ollama/api"
)

const requestTimeout = 30 * time.Second

// OllamaProvider is a implementation of AIProvider interface for Ollama
type OllamaProvider struct {
	baseURL         string
	model           string
	client          *api.Client
	maxDiffSize     int
	promptGenerator services.PromptGenerator
	log             utils.Logger
}

// NewOllamaProvider creates a new instance of OllamaProvider
func NewOllamaProvider(
	baseURL string,
	model string,
	maxDiffSize int,
	promptGenerator services.PromptGenerator,
	log utils.Logger,
) (*OllamaProvider, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	client := api.NewClient(parsedURL, httpClient)
	return &OllamaProvider{
		baseURL:         baseURL,
		model:           model,
		client:          client,
		maxDiffSize:     maxDiffSize,
		promptGenerator: promptGenerator,
		log:             log,
	}, nil
}

func (p *OllamaProvider) GenerateCommitMessage(diff string, changedFiles []string) (string, error) {
	prompt := p.promptGenerator.GenerateCommitPrompt(diff, changedFiles, p.maxDiffSize)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	p.log.Debug("Generating commit message with Ollama: %s", prompt)

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

	p.log.Debug("AI response commit message: %s", commitMessage)

	return strings.TrimSpace(commitMessage), nil
}

// Vérification statique que OllamaProvider implémente bien l'interface AIProvider
var _ services.AIProvider = (*OllamaProvider)(nil)
