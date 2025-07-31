package server

import (
	"log"
	"os"
	"strings"

	"instagram-downloader-api/internal/handlers"
	"instagram-downloader-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

// New creates and configures a new Gin server
func New() *gin.Engine {
	// Set Gin mode (will be "release" in production)
	gin.SetMode(gin.DebugMode)

	// Create router
	router := gin.New()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())

	// --- Load Instagram Session IDs from environment variable ---
	// It's recommended to store your session IDs in an environment variable,
	// separated by commas.
	// Example: INSTAGRAM_SESSION_IDS="sessionid1,sessionid2,sessionid3"
	sessionIDsStr := os.Getenv("INSTAGRAM_SESSION_IDS")
	if sessionIDsStr == "" {
		log.Println("WARNING: INSTAGRAM_SESSION_IDS environment variable not set. Requests to Instagram will be anonymous and are likely to be rate-limited.")
	}
	sessionIDs := strings.Split(sessionIDsStr, ",")
	// Trim whitespace from each session ID
	for i, id := range sessionIDs {
		sessionIDs[i] = strings.TrimSpace(id)
	}
	// ---

	// Create handlers
	instagramHandler := handlers.NewInstagramHandler(handlers.InstagramHandlerConfig{
		SessionIDs: sessionIDs,
	})
	downloadHandler := handlers.NewDownloadHandler()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Instagram Downloader API is running",
		})
	})

	// API routes
	api := router.Group("/api")
	{
		// Instagram routes
		instagram := api.Group("/instagram")
		{
			instagram.GET("/p/:shortcode", instagramHandler.GetInstagramPost)
		}

		// Download proxy route
		api.GET("/download-proxy", downloadHandler.DownloadProxy)
	}

	return router
}
