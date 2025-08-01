package server

import (
	"log"
	"os"
	"strings"

	"instagram-downloader-api/internal/handlers"
	"instagram-downloader-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	sessionIDsStr := os.Getenv("INSTAGRAM_SESSION_IDS")
	if sessionIDsStr == "" {
		log.Println("WARNING: INSTAGRAM_SESSION_IDS environment variable not set. Requests to Instagram will be anonymous and are likely to be rate-limited.")
	}
	sessionIDs := strings.Split(sessionIDsStr, ",")
	for i, id := range sessionIDs {
		sessionIDs[i] = strings.TrimSpace(id)
	}
	instagramHandler := handlers.NewInstagramHandler(handlers.InstagramHandlerConfig{
		SessionIDs: sessionIDs,
	})
	downloadHandler := handlers.NewDownloadHandler()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Instagram Downloader API is running",
		})
	})
	api := router.Group("/api")
	{
		instagram := api.Group("/instagram")
		{
			instagram.GET("/p/:shortcode", instagramHandler.GetInstagramPost)
		}
		api.GET("/download-proxy", downloadHandler.DownloadProxy)
	}
	return router
}
