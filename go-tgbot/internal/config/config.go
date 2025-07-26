package config

import (
	"encoding/json"
	"go-tgbot/internal/types"
	"os"
)

func Load() (*types.Config, error) {
	var cfg types.Config

	// For now, we'll use a placeholder.
	// In a real application, you might load this from a file or environment variables.
	cfg.TelegramBotToken = "8222315138:AAGSPdH-95kLR7N1mMnVj-oQB03ymiD6eRI" // Replace with your actual test key

	return &cfg, nil
}

// Save is a placeholder for saving configuration, not used in this example
func Save(cfg *types.Config) error {
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("config.json", data, 0644)
}
