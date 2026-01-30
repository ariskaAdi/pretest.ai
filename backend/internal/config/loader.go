package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    GoogleAIAPIKey string
    Port           string
    Environment    string
}

func Load() (*Config, error) {
    if err := godotenv.Load(); err != nil {
        godotenv.Load("../../.env")
    }
    
    apiKey := os.Getenv("GOOGLE_AI_API_KEY")
    if apiKey == "" {
        return nil, fmt.Errorf("GOOGLE_AI_API_KEY is required")
    }
    
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    
    env := os.Getenv("ENV")
    if env == "" {
        env = "development"
    }
    
    return &Config{
        GoogleAIAPIKey: apiKey,
        Port:           port,
        Environment:    env,
    }, nil
}