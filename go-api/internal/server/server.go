package server

import (
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

	// Create handlers
	instagramHandler := handlers.NewInstagramHandler()
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
