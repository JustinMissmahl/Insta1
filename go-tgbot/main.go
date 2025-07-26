package main

import (
	"context"
	"go-tgbot/internal/config"
	"go-tgbot/internal/handler"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
)

// Send any text message to the bot after the bot has been started

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	b, err := bot.New(cfg.TelegramBotToken)
	if err != nil {
		panic(err)
	}

	botHandler := handler.New(cfg)
	botHandler.Register(b)

	log.Println("Bot started...")
	b.Start(ctx)
}
