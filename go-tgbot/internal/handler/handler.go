package handler

import (
	"bytes"
	"context"
	"fmt"
	"go-tgbot/internal/api"
	"go-tgbot/internal/types"
	"go-tgbot/internal/util"
	"io"
	"log"
	"os"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

const (
	downloadPath = "downloads"
)

type BotHandler struct {
	cfg       *types.Config
	apiClient *api.Client
}

func New(cfg *types.Config) *BotHandler {
	return &BotHandler{
		cfg:       cfg,
		apiClient: api.New(cfg.ApiBaseURL),
	}
}

func (h *BotHandler) Register(b *bot.Bot) {
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, h.handleMessage)
}

func (h *BotHandler) handleMessage(ctx context.Context, b *bot.Bot, update *models.Update) {
	log.Printf("Received message from %s: %s", update.Message.From.Username, update.Message.Text)

	shortcode := util.ExtractShortcode(update.Message.Text)
	if shortcode == "" {
		log.Printf("Invalid URL: %s", update.Message.Text)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Please send a valid Instagram post or reel URL.",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Processing your request...",
	})

	postData, err := h.apiClient.GetInstagramPostData(shortcode)
	if err != nil {
		log.Printf("Failed to get post data for shortcode %s: %v", shortcode, err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Failed to get post data. The post may be private or deleted.",
		})
		return
	}

	if !postData.IsVideo {
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "This post is not a video.",
		})
		return
	}

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Downloading video by %s...", postData.Owner.Username),
	})

	filePath, err := h.apiClient.DownloadVideo(postData.VideoURL, shortcode, downloadPath)
	if err != nil {
		log.Printf("Failed to download video for shortcode %s: %v", shortcode, err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Failed to download video.",
		})
		return
	}
	defer os.Remove(filePath) // Defer deletion of the file

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Download complete. Sending video...",
	})

	caption := fmt.Sprintf("Video by %s (@%s)", postData.Owner.FullName, postData.Owner.Username)

	videoFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open video file %s: %v", filePath, err)
		return
	}
	defer videoFile.Close()

	videoData, err := io.ReadAll(videoFile)
	if err != nil {
		log.Printf("Failed to read video file %s: %v", filePath, err)
		return
	}

	_, err = b.SendVideo(ctx, &bot.SendVideoParams{
		ChatID:  update.Message.Chat.ID,
		Video:   &models.InputFileUpload{Filename: fmt.Sprintf("%s.mp4", shortcode), Data: bytes.NewReader(videoData)},
		Caption: caption,
	})
	if err != nil {
		log.Printf("Failed to send video for shortcode %s: %v", shortcode, err)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Failed to send video.",
		})
		return
	}

	log.Printf("Successfully sent video for shortcode %s", shortcode)
}
