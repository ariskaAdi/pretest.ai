package database

import (
	"ariskaAdi-pretest-ai/internal/config"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	if err := godotenv.Load("../../.env"); err != nil {
	log.Println(".env not found, using OS env")
	}

	config.LoadConfig()
}

func TestConnectionPostgres(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})
}