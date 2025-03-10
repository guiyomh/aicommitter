package ai

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
	"github.com/ollama/ollama/api"
)

const (
	requestTimeout         = 30 * time.Second
	ollamaTemperature      = 0.7
	ollamaTopK             = 40
	ollamaTopP             = 0.9
	ollamaNumPredict       = 128
	ollamaNumCtx           = 4096
	ollamaRepeatPenalty    = 1.1
	ollamaFrequencyPenalty = 1.1
	ollamaPresencePenalty  = 0.0
)

// OllamaProvider is a implementation of AIProvider interface for Ollama
type OllamaProvider struct {
	baseURL            string
	model              string
	client             *api.Client
	maxDiffSize        int
	promptGenerator    services.PromptGenerator
	log                utils.Logger
	diffFormatter      services.DiffFormatter
	commitTypeProvider services.CommitTypeProvider
}

// NewOllamaProvider creates a new instance of OllamaProvider
func NewOllamaProvider(
	baseURL string,
	model string,
	maxDiffSize int,
	promptGenerator services.PromptGenerator,
	log utils.Logger,
	diffFormatter services.DiffFormatter,
	commitTypeProvider services.CommitTypeProvider,
) (*OllamaProvider, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	httpClient := &http.Client{}
	client := api.NewClient(parsedURL, httpClient)
	return &OllamaProvider{
		baseURL:            baseURL,
		model:              model,
		client:             client,
		maxDiffSize:        maxDiffSize,
		promptGenerator:    promptGenerator,
		log:                log,
		diffFormatter:      diffFormatter,
		commitTypeProvider: commitTypeProvider,
	}, nil
}

func (p *OllamaProvider) GenerateCommitMessage(diff string, changedFiles []string) (services.AIResponse, error) {
	systemPrompt := p.promptGenerator.GenerateCommitPrompt(p.maxDiffSize)

	formattedDiff := p.diffFormatter.FormatDiff(diff, p.maxDiffSize)
	fileStr := p.diffFormatter.GetFileList(changedFiles)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	p.log.Debug("Generating commit message with Ollama: %s", systemPrompt)

	var commitMessage string

	respFunc := func(resp api.GenerateResponse) error {
		commitMessage += resp.Response
		return nil
	}

	prompt := fmt.Sprintf("%s\n\n%s", formattedDiff, fileStr)

	req := &api.GenerateRequest{
		Model:  p.model,
		System: systemPrompt,
		Prompt: prompt,
		Stream: new(bool),
		Options: map[string]interface{}{
			"temperature":       ollamaTemperature,
			"top_k":             ollamaTopK,
			"top_p":             ollamaTopP,
			"num_predict":       ollamaNumPredict,
			"num_ctx":           ollamaNumCtx,
			"repeat_penalty":    ollamaRepeatPenalty,
			"frequency_penalty": ollamaFrequencyPenalty,
			"presence_penalty":  ollamaPresencePenalty,
		},
	}

	err := p.client.Generate(ctx, req, respFunc)
	if err != nil {
		return "", fmt.Errorf("error while generating commit message with Ollama: %w", err)
	}

	p.log.Debug("AI response commit message: %s", commitMessage)

	return services.AIResponse(commitMessage), nil
}

// Vérification statique que OllamaProvider implémente bien l'interface AIProvider
var _ services.AIProvider = (*OllamaProvider)(nil)
