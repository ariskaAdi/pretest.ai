package main

import (
	"ariskaAdi-pretest-ai/apps/quiz"
	"ariskaAdi-pretest-ai/external/database"
	"ariskaAdi-pretest-ai/internal/config"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load("../../.env"); err != nil {
	log.Println(".env not found, using OS env")
	}

	config.LoadConfig()

	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}

	if db != nil {
		log.Println("DB CONNECTED")
	}

	g := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{
			APIKey: config.Cfg.Genkit.GoogleAIAPIKey,
		}),
		genkit.WithDefaultModel("googleai/gemini-2.5-flash"),
	)

	app := fiber.New(fiber.Config{
		AppName: config.Cfg.App.Name,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTION",
		AllowHeaders: "Authorization, Content-Type",
		AllowCredentials: true,
	}))


	quiz.Init(app, g)

	go func() {
		if err := app.Listen(":" + config.Cfg.App.Port); err != nil {
			log.Println("Server stopped:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down server...")

	if err := app.Shutdown(); err != nil {
		log.Println("Shutdown error:", err)
	}

	log.Println("Server exited gracefully")
}


