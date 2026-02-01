package quiz

import (
	"github.com/firebase/genkit/go/genkit"
	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router, genkit *genkit.Genkit) {
	svc := NewService(genkit)

	handler := newHandler(svc)

	quizRoute := router.Group("/quiz")
	{
		quizRoute.Post("/generate", handler.GenerateQuiz)
		quizRoute.Post("/pdf-local", handler.GenerateQuizFromPdfLocal)
	}
}