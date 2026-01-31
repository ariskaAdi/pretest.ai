package main

import (
	"ariskaAdi-pretest-ai/apps/quiz"
	"ariskaAdi-pretest-ai/internal/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
	log.Println(".env not found, using OS env")
}

	config.LoadConfig()

	app := fiber.New(fiber.Config{
		AppName: config.Cfg.App.Name,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET,POST,PUT,PATCH,DELETE,OPTION",
		AllowHeaders: "Authorization, Content-Type",
		AllowCredentials: true,
	}))


	quiz.Init(app, &config.Cfg)

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


