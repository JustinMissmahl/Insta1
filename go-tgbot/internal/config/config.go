package config

import (
	"go-tgbot/internal/types"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

func Load() (*types.Config, error) {
	godotenv.Load()
	cfg := &types.Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		ApiBaseURL:       os.Getenv("API_BASE_URL"),
	}
	if cfg.ApiBaseURL == "" {
		cfg.ApiBaseURL = "http://localhost:8080"
	}
	allowedUserIDsStr := os.Getenv("ALLOWED_USER_IDS")
	if allowedUserIDsStr != "" {
		idStrings := strings.Split(allowedUserIDsStr, ",")
		allowedIDs := make([]int64, 0, len(idStrings))
		for _, idStr := range idStrings {
			id, err := strconv.ParseInt(strings.TrimSpace(idStr), 10, 64)
			if err != nil {
				continue
			}
			allowedIDs = append(allowedIDs, id)
		}
		cfg.AllowedUserIDs = allowedIDs
	}
	return cfg, nil
}
