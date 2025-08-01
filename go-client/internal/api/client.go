package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"instagram-api-test-client/internal/types"
)

const (
	APIBaseURL = "http://localhost:8080"
)

func GetInstagramPostData(shortcode string) (*types.InstagramResponse, error) {
	url := fmt.Sprintf("%s/api/instagram/p/%s", APIBaseURL, shortcode)
	log.Printf("🔍 Fetching Instagram post data from: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch post data: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}
	var igResponse types.InstagramResponse
	if err := json.NewDecoder(resp.Body).Decode(&igResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	return &igResponse, nil
}

func DownloadVideo(videoURL, filename string) error {
	encodedURL := url.QueryEscape(videoURL)
	downloadURL := fmt.Sprintf("%s/api/download-proxy?url=%s&filename=%s", APIBaseURL, encodedURL, filename)
	log.Printf("⬇️  Downloading video from: %s", downloadURL)
	resp, err := http.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download video: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status: %d", resp.StatusCode)
	}
	downloadsDir := "downloads"
	if err := os.MkdirAll(downloadsDir, 0755); err != nil {
		return fmt.Errorf("failed to create downloads directory: %w", err)
	}
	filePath := filepath.Join(downloadsDir, filename)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	written, err := io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write video content: %w", err)
	}
	log.Printf("✅ Video downloaded successfully!")
	log.Printf("📁 File: %s", filePath)
	log.Printf("📊 Size: %.2f MB", float64(written)/(1024*1024))
	return nil
}

func TestHealthCheck() error {
	log.Printf("🏥 Testing health check...")
	resp, err := http.Get(fmt.Sprintf("%s/health", APIBaseURL))
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health check returned status: %d", resp.StatusCode)
	}
	var healthResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&healthResponse); err != nil {
		return fmt.Errorf("failed to decode health response: %w", err)
	}
	log.Printf("✅ Health check passed: %v", healthResponse["message"])
	return nil
}
