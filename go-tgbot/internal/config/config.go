package config

import (
	"go-tgbot/internal/types"
	"os"

	"github.com/joho/godotenv"
)

func Load() (*types.Config, error) {
	// Load .env file if it exists
	godotenv.Load()

	cfg := &types.Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		ApiBaseURL:       os.Getenv("API_BASE_URL"),
	}

	if cfg.ApiBaseURL == "" {
		cfg.ApiBaseURL = "http://localhost:8080"
	}

	return cfg, nil
}
