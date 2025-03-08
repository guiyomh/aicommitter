package ai

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/guiyomh/aicommitter/internal/domain/services"
	"github.com/guiyomh/aicommitter/internal/domain/utils"
	"github.com/ollama/ollama/api"
)

const (
	requestTimeout         = 30 * time.Second
	ollamaTemperature      = 0.8
	ollamaTopK             = 40
	ollamaTopP             = 0.9
	ollamaNumPredict       = 128
	ollamaRepeatPenalty    = 1.1
	ollamaFrequencyPenalty = 1.1
	ollamaPresencePenalty  = 0.0
)

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
		Options: map[string]interface{}{
			"temperature":       ollamaTemperature,
			"top_k":             ollamaTopK,
			"top_p":             ollamaTopP,
			"num_predict":       ollamaNumPredict,
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

	// Extraire le JSON entre les délimiteurs
	start := strings.Index(commitMessage, "---COMMIT_MESSAGE_START---")
	end := strings.Index(commitMessage, "---COMMIT_MESSAGE_END---")

	if start == -1 || end == -1 {
		return "", fmt.Errorf("délimiteur de début ou de fin non trouvé dans la réponse")
	}

	jsonStr := strings.TrimSpace(commitMessage[start+len("---COMMIT_MESSAGE_START---") : end])

	// Parser le JSON
	var commit struct {
		Type        string `json:"type"`
		Scope       string `json:"scope"`
		Description string `json:"description"`
		Body        string `json:"body"`
		Footer      string `json:"footer"`
	}

	if err := json.Unmarshal([]byte(jsonStr), &commit); err != nil {
		return "", fmt.Errorf("erreur lors du parsing du JSON: %w", err)
	}

	// Construire le message de commit conventionnel
	message := commit.Type
	const nullValue = "null"
	if commit.Scope != nullValue && commit.Scope != "" {
		message += "(" + commit.Scope + ")"
	}
	message += ": " + commit.Description

	if commit.Body != nullValue && commit.Body != "" {
		message += "\n\n" + commit.Body
	}

	if commit.Footer != nullValue && commit.Footer != "" {
		message += "\n\n" + commit.Footer
	}

	return message, nil
}

// Vérification statique que OllamaProvider implémente bien l'interface AIProvider
var _ services.AIProvider = (*OllamaProvider)(nil)
