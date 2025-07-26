package main

import (
	"fmt"
	"os"
	"testing"
	"time"

	"instagram-api-test-client/internal/api"
	"instagram-api-test-client/internal/utils"
)

func TestDownloadFlow(t *testing.T) {
	// Set a mock environment or use a test server if available.
	// For now, we are hitting the live development API.
	if os.Getenv("CI") != "" {
		t.Skip("Skipping test that hits a live API in CI environment")
	}

	instagramURL := "https://www.instagram.com/reel/DJeXBKNPFNM/?igsh=aTV1ajNmNXhtMDB3"

	// Step 1: Health Check
	if err := api.TestHealthCheck(); err != nil {
		t.Fatalf("Health check failed: %v", err)
	}

	// Step 2: Extract Shortcode
	shortcode := utils.ExtractShortcode(instagramURL)
	if shortcode == "" {
		t.Fatalf("Failed to extract shortcode from URL: %s", instagramURL)
	}
	t.Logf("Extracted shortcode: %s", shortcode)

	// Step 3: Get Post Data
	postData, err := api.GetInstagramPostData(shortcode)
	if err != nil {
		t.Fatalf("Failed to get post data: %v", err)
	}

	media := postData.Data.XdtShortcodeMedia
	if !media.IsVideo {
		t.Log("Post is not a video, skipping download.")
		return
	}

	// Step 4: Download Video
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("instagram_video_%s_%s.mp4", shortcode, timestamp)
	if err := api.DownloadVideo(media.VideoURL, filename); err != nil {
		t.Fatalf("Failed to download video: %v", err)
	}

	t.Log("Test completed successfully!")
}
