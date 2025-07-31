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

	var filePath string
	var mediaType string
	var downloadErr error

	if postData.IsVideo {
		mediaType = "video"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Downloading video by %s...", postData.Owner.Username),
		})
		filePath, downloadErr = h.apiClient.DownloadVideo(postData.VideoURL, shortcode, downloadPath)
	} else {
		mediaType = "image"
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Downloading image by %s...", postData.Owner.Username),
		})
		// Get the highest resolution image
		imageURL := postData.DisplayUrl
		if len(postData.DisplayResources) > 0 {
			imageURL = postData.DisplayResources[len(postData.DisplayResources)-1].Src
		}
		filePath, downloadErr = h.apiClient.DownloadImage(imageURL, shortcode, downloadPath)
	}

	if downloadErr != nil {
		log.Printf("Failed to download %s for shortcode %s: %v", mediaType, shortcode, downloadErr)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Failed to download %s.", mediaType),
		})
		return
	}
	defer os.Remove(filePath) // Defer deletion of the file

	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Download complete. Sending %s...", mediaType),
	})

	caption := fmt.Sprintf("%s by %s (@%s)", mediaType, postData.Owner.FullName, postData.Owner.Username)

	mediaFile, err := os.Open(filePath)
	if err != nil {
		log.Printf("Failed to open %s file %s: %v", mediaType, filePath, err)
		return
	}
	defer mediaFile.Close()

	mediaData, err := io.ReadAll(mediaFile)
	if err != nil {
		log.Printf("Failed to read %s file %s: %v", mediaType, filePath, err)
		return
	}

	var sendErr error
	if postData.IsVideo {
		_, sendErr = b.SendVideo(ctx, &bot.SendVideoParams{
			ChatID:  update.Message.Chat.ID,
			Video:   &models.InputFileUpload{Filename: fmt.Sprintf("%s.mp4", shortcode), Data: bytes.NewReader(mediaData)},
			Caption: caption,
		})
	} else {
		_, sendErr = b.SendPhoto(ctx, &bot.SendPhotoParams{
			ChatID:  update.Message.Chat.ID,
			Photo:   &models.InputFileUpload{Filename: fmt.Sprintf("%s.jpg", shortcode), Data: bytes.NewReader(mediaData)},
			Caption: caption,
		})
	}

	if sendErr != nil {
		log.Printf("Failed to send %s for shortcode %s: %v", mediaType, shortcode, sendErr)
		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   fmt.Sprintf("Failed to send %s.", mediaType),
		})
		return
	}

	log.Printf("Successfully sent %s for shortcode %s", mediaType, shortcode)
}
