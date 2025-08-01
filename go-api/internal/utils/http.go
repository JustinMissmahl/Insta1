package utils

import (
	"instagram-downloader-api/internal/types"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RespondWithError(c *gin.Context, statusCode int, errorType, message string) {
	errorResponse := types.ErrorResponse{
		Error:      errorType,
		Message:    message,
		StatusCode: types.HTTPCode(statusCode),
	}
	c.JSON(statusCode, errorResponse)
}

func RespondWithData(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}

func RespondWithJSON(c *gin.Context, statusCode int, payload interface{}) {
	c.JSON(statusCode, payload)
}

func SetDownloadHeaders(c *gin.Context, filename, contentType string, contentLength int64) {
	c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	c.Header("Content-Type", contentType)
	if contentLength > 0 {
		c.Header("Content-Length", strconv.FormatInt(contentLength, 10))
	}
}
