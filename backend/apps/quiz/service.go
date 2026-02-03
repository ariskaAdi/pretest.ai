package quiz

import (
	infrarequest "ariskaAdi-pretest-ai/infra/request"
	"context"
	"encoding/base64"
	"fmt"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
)

type service struct {
	genkit *genkit.Genkit
}

func NewService(g *genkit.Genkit) *service {
	return &service{genkit: g}
}

func (s *service) GenerateQuiz(ctx context.Context, req GenerateQuizRequest) (*QuizResultWithStats, error) {


	// pdfBytes, err := fetchPdf(ctx, req.PdfUrl)
	// if err != nil {
	// 	return nil, err
	// }

	// pdfBase64 := base64.StdEncoding.EncodeToString(pdfBytes)

	pdfPart := ai.NewMediaPart("application/pdf", req.PdfUrl)
    
	promptPart := ai.NewTextPart(infrarequest.GenerateQuizPrompt)

	message := &ai.Message{
		Role: ai.RoleUser,
		Content:  []*ai.Part{pdfPart, promptPart},
	}

	resp, meta, err := genkit.GenerateData[QuizResponse](
		ctx,
		s.genkit,
		ai.WithMessages(message),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate quiz: %w", err)
	}

	result := &QuizResultWithStats{
		Quiz: resp,
	}

	if meta != nil && meta.Usage != nil {
        result.Usage = UsageInfo{
            InputTokens:  meta.Usage.InputTokens,
            OutputTokens: meta.Usage.OutputTokens,
            TotalTokens:  meta.Usage.TotalTokens,
        }
    }

    return result, nil
}

func (s *service) GenerateQuizFromBytes(ctx context.Context, data []byte) (*QuizResponse, error) {
    base64Data := base64.StdEncoding.EncodeToString(data)
    dataURI := fmt.Sprintf("data:application/pdf;base64,%s", base64Data)

    message := &ai.Message{
        Role: ai.RoleUser,
        Content: []*ai.Part{
            ai.NewMediaPart("application/pdf", dataURI),
            ai.NewTextPart(infrarequest.GenerateQuizPrompt),
        },
    }

    resp, _, err := genkit.GenerateData[QuizResponse](ctx, s.genkit, ai.WithMessages(message))
    return resp, err
}