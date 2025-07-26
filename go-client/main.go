package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"instagram-api-test-client/internal/api"
	"instagram-api-test-client/internal/utils"
)

func main() {
	fmt.Println("🚀 Instagram Downloader Client")
	fmt.Println("=======================================")

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <instagram_url>")
		os.Exit(1)
	}
	instagramURL := os.Args[1]

	// Step 1: Test health check
	if err := api.TestHealthCheck(); err != nil {
		log.Fatalf("❌ Health check failed: %v", err)
	}

	// Step 2: Extract shortcode
	shortcode := utils.ExtractShortcode(instagramURL)
	if shortcode == "" {
		log.Fatalf("❌ Failed to extract shortcode from URL: %s", instagramURL)
	}

	log.Printf("📝 Extracted shortcode: %s", shortcode)

	// Step 3: Get post data
	log.Printf("🔗 Processing Instagram URL: %s", instagramURL)

	postData, err := api.GetInstagramPostData(shortcode)
	if err != nil {
		log.Fatalf("❌ Failed to get post data: %v", err)
	}

	media := postData.Data.XdtShortcodeMedia

	// Step 4: Display post information
	fmt.Println("\n📊 Post Information:")
	fmt.Printf("  ID: %s\n", media.ID)
	fmt.Printf("  Is Video: %t\n", media.IsVideo)
	fmt.Printf("  Duration: %.1f seconds\n", media.VideoDuration)
	fmt.Printf("  Owner: %s (@%s)\n", media.Owner.FullName, media.Owner.Username)
	fmt.Printf("  Video URL: %s\n", media.VideoURL)

	if !media.IsVideo {
		log.Println("✅ Post is not a video. Nothing to download.")
		return
	}

	// Step 5: Download the video
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("instagram_video_%s_%s.mp4", shortcode, timestamp)

	fmt.Printf("\n⬇️  Starting download...\n")

	if err := api.DownloadVideo(media.VideoURL, filename); err != nil {
		log.Fatalf("❌ Failed to download video: %v", err)
	}

	fmt.Println("\n🎉 Download completed successfully!")
	fmt.Println("📁 Check the 'downloads' folder for your video.")
}
