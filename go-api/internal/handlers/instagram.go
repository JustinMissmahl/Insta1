package handlers

import (
	"net/http"

	"instagram-downloader-api/internal/instagram"
	"instagram-downloader-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type InstagramHandler struct {
	igClient *instagram.Client
}

type InstagramHandlerConfig struct {
	SessionIDs []string
}

func NewInstagramHandler(config InstagramHandlerConfig) *InstagramHandler {
	return &InstagramHandler{
		igClient: instagram.NewClient(config.SessionIDs),
	}
}

func (h *InstagramHandler) GetInstagramPost(c *gin.Context) {
	shortcode := c.Param("shortcode")
	if shortcode == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "noShortcode", "shortcode is required")
		return
	}
	response, statusCode, err := h.igClient.GetPostGraphQL(shortcode)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "serverError", err.Error())
		return
	}
	switch statusCode {
	case http.StatusOK:
		if response.Data.XdtShortcodeMedia.ID == "" {
			utils.RespondWithError(c, http.StatusNotFound, "notFound", "post not found")
			return
		}
		utils.RespondWithData(c, http.StatusOK, response.Data)
		return
	case http.StatusNotFound:
		utils.RespondWithError(c, http.StatusNotFound, "notFound", "post not found")
		return
	case http.StatusTooManyRequests, http.StatusUnauthorized:
		utils.RespondWithError(c, http.StatusTooManyRequests, "tooManyRequests", "too many requests, try again later")
		return
	default:
		utils.RespondWithError(c, http.StatusInternalServerError, "serverError", "Failed to fetch post data")
		return
	}
}
