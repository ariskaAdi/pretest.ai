package quiz

import (
	"ariskaAdi-pretest-ai/internal/config"
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Init(router fiber.Router, cfg *config.Config) {
	ctx := context.Background()

	// Init service dengan context
	svc, err := NewService(ctx, cfg)
	if err != nil {
		log.Fatal("Failed to init quiz service:", err)
	}

	handler := newHandler(svc)

	quizRoute := router.Group("/quiz")
	{
		quizRoute.Post("/", handler.GenerateQuiz)
	}
}