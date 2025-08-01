package handlers

import (
	"io"
	"net/http"
	"strconv"
	"strings"

	"instagram-downloader-api/internal/types"
	"instagram-downloader-api/internal/utils"

	"github.com/gin-gonic/gin"
)

type DownloadHandler struct {
	httpClient *http.Client
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{
		httpClient: &http.Client{},
	}
}

func (h *DownloadHandler) DownloadProxy(c *gin.Context) {
	var req types.DownloadProxyRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "invalidParams", err.Error())
		return
	}
	if req.URL == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "missingUrl", "url is required")
		return
	}
	if req.Filename == "" {
		req.Filename = "gram-grabberz-video.mp4"
	}
	if !strings.HasPrefix(req.URL, "https://") {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid URL format", "URL must start with https://")
		return
	}
	resp, err := h.httpClient.Get(req.URL)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "serverError", "Failed to fetch video: "+err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		utils.RespondWithError(c, http.StatusInternalServerError, "serverError", "Failed to fetch video: "+resp.Status)
		return
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		contentType = "video/mp4"
	}
	contentLengthStr := resp.Header.Get("Content-Length")
	var contentLength int64
	if contentLengthStr != "" {
		if parsed, err := strconv.ParseInt(contentLengthStr, 10, 64); err == nil {
			contentLength = parsed
		}
	}
	utils.SetDownloadHeaders(c, req.Filename, contentType, contentLength)
	c.Status(http.StatusOK)
	_, err = io.Copy(c.Writer, resp.Body)
	if err != nil {
		c.Error(err)
		return
	}
}
