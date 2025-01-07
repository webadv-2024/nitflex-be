package llm_search

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type LLMSearchAdapter interface {
	SearchMoviesLLM(ctx context.Context, query string) ([]string, error)
}

type llmSearchAdapter struct {
	openAIClient *openai.Client
}

type Config struct {
	OpenAIApiKey string
}

func NewLLMSearchAdapter(cfg Config) LLMSearchAdapter {
	return &llmSearchAdapter{
		openAIClient: openai.NewClient(cfg.OpenAIApiKey),
	}
}

// SearchMovies implements the LLMSearchAdapter interface
func (llm *llmSearchAdapter) SearchMoviesLLM(ctx context.Context, query string) ([]string, error) {
	resp, err := llm.openAIClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: `You are a movie search assistant. Given a search query, return relevant movies that match the query. The title should be match with imdb title.
					The response only contains [title1, title2, ...]`,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		return nil, err
	}

	fmt.Println(resp.Choices[0].Message.Content)

	return []string{resp.Choices[0].Message.Content}, nil
}