package handlers

import (
	"log"
	"net/http"

	"instagram-downloader-api/internal/instagram"
	"instagram-downloader-api/internal/utils"

	"github.com/gin-gonic/gin"
)

// InstagramHandler handles Instagram-related requests
type InstagramHandler struct {
	igClient *instagram.Client
}

// InstagramHandlerConfig holds the dependencies for the Instagram handler.
type InstagramHandlerConfig struct {
	SessionIDs []string
}

// NewInstagramHandler creates a new Instagram handler
func NewInstagramHandler(config InstagramHandlerConfig) *InstagramHandler {
	return &InstagramHandler{
		igClient: instagram.NewClient(config.SessionIDs),
	}
}

// GetInstagramPost handles GET /api/instagram/p/:shortcode
func (h *InstagramHandler) GetInstagramPost(c *gin.Context) {
	shortcode := c.Param("shortcode")

	if shortcode == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "noShortcode", "shortcode is required")
		return
	}

	response, statusCode, err := h.igClient.GetPostGraphQL(shortcode)
	if err != nil {
		log.Printf("Instagram API error: %v", err)
		utils.RespondWithError(c, http.StatusInternalServerError, "serverError", err.Error())
		return
	}

	switch statusCode {
	case http.StatusOK:
		// Check if data exists
		if response.Data.XdtShortcodeMedia.ID == "" {
			utils.RespondWithError(c, http.StatusNotFound, "notFound", "post not found")
			return
		}

		// Check if it's a video
		if !response.Data.XdtShortcodeMedia.IsVideo {
			utils.RespondWithError(c, http.StatusBadRequest, "notVideo", "post is not a video")
			return
		}

		// Return successful response matching Next.js structure
		utils.RespondWithData(c, http.StatusOK, response.Data)
		return

	case http.StatusNotFound:
		utils.RespondWithError(c, http.StatusNotFound, "notFound", "post not found")
		return

	case http.StatusTooManyRequests, http.StatusUnauthorized:
		utils.RespondWithError(c, http.StatusTooManyRequests, "tooManyRequests", "too many requests, try again later")
		return

	default:
		log.Printf("Unexpected status code from Instagram API: %d", statusCode)
		utils.RespondWithError(c, http.StatusInternalServerError, "serverError", "Failed to fetch post data")
		return
	}
}
