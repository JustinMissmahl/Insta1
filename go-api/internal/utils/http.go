package utils

import (
	"strconv"

	"instagram-downloader-api/internal/types"

	"github.com/gin-gonic/gin"
)

// RespondWithError sends an error response with the specified status code
func RespondWithError(c *gin.Context, statusCode int, errorType, message string) {
	errorResponse := types.ErrorResponse{
		Error:      errorType,
		Message:    message,
		StatusCode: types.HTTPCode(statusCode),
	}

	c.JSON(statusCode, errorResponse)
}

// RespondWithData sends a successful response with data
func RespondWithData(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, gin.H{"data": data})
}

// RespondWithJSON sends a JSON response
func RespondWithJSON(c *gin.Context, statusCode int, payload interface{}) {
	c.JSON(statusCode, payload)
}

// SetDownloadHeaders sets headers for file download responses
func SetDownloadHeaders(c *gin.Context, filename, contentType string, contentLength int64) {
	c.Header("Content-Disposition", "attachment; filename=\""+filename+"\"")
	c.Header("Content-Type", contentType)

	if contentLength > 0 {
		c.Header("Content-Length", strconv.FormatInt(contentLength, 10))
	}
}
