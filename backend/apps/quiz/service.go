package quiz

import (
	infrarequest "ariskaAdi-pretest-ai/infra/request"
	"ariskaAdi-pretest-ai/internal/config"
	"context"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

type service struct {
	genkit *genkit.Genkit
}

// NewService - init service dengan genkit
func NewService(ctx context.Context, cfg *config.Config) (*service, error) {
	g := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{
			APIKey: cfg.GoogleAIAPIKey,
		}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	return &service{genkit: g}, nil
}

func (s *service) GenerateQuiz(ctx context.Context, text string) (*QuizResponse, error) {
    
	prompt := fmt.Sprintf(infrarequest.GenerateQuizPrompt, text)

	resp, _, err := genkit.GenerateData[QuizResponse](
		ctx,
		s.genkit,
		ai.WithPrompt(prompt),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate quiz: %w", err)
	}

	return resp, nil
}