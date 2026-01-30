package main

import (
	"ariskaAdi-pretest-ai/apps/quiz"
	"ariskaAdi-pretest-ai/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Setup Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
	})

	// Middleware~
	app.Use(logger.New())
	app.Use(cors.New())

	// Health check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"env":    cfg.Environment,
		})
	})

	// Routes
	api := app.Group("/api")
	quiz.Init(api, cfg)

	// Start server
	log.Printf("🚀 Server running on port %s (env: %s)", cfg.Port, cfg.Environment)
	log.Fatal(app.Listen(":" + cfg.Port))
}


