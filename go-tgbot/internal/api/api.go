package api

import (
	"encoding/json"
	"fmt"
	"go-tgbot/internal/types"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type Client struct {
	baseURL    string
	httpClient *http.Client
}

func New(apiBaseURL string) *Client {
	return &Client{
		baseURL:    apiBaseURL,
		httpClient: &http.Client{},
	}
}

func (c *Client) GetInstagramPostData(shortcode string) (*types.XdtShortcodeMediaDto, error) {
	url := fmt.Sprintf("%s/api/instagram/p/%s", c.baseURL, shortcode)

	var lastErr error

	maxRetries := 4
	retryDelay := 1 * time.Second

	for i := 0; i < maxRetries; i++ {
		req, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("failed to execute request: %w", err)
			log.Printf("Attempt %d: failed to execute request for %s: %v. Retrying in %v...", i+1, shortcode, err, retryDelay)
			time.Sleep(retryDelay)
			retryDelay *= 2
			continue
		}

		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("failed to read response body: %w", err)
			}

			var response types.IGGraphQLResponseDto
			if err := json.Unmarshal(body, &response); err != nil {
				return nil, fmt.Errorf("failed to unmarshal response: %w", err)
			}

			return &response.Data.XdtShortcodeMedia, nil
		}

		bodyBytes, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == http.StatusTooManyRequests {
			lastErr = fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
			log.Printf("Attempt %d: Rate limit hit for shortcode %s. Retrying in %v...", i+1, shortcode, retryDelay)
			time.Sleep(retryDelay)
			retryDelay *= 2
			continue
		}

		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil, fmt.Errorf("failed after %d retries for shortcode %s: %w", maxRetries, shortcode, lastErr)
}

func (c *Client) DownloadVideo(videoURL, shortcode, downloadPath string) (string, error) {
	if _, err := os.Stat(downloadPath); os.IsNotExist(err) {
		if err := os.MkdirAll(downloadPath, os.ModePerm); err != nil {
			return "", fmt.Errorf("failed to create download directory: %w", err)
		}
	}

	filePath := filepath.Join(downloadPath, fmt.Sprintf("%s.mp4", shortcode))

	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	var lastErr error
	maxRetries := 3
	retryDelay := 1 * time.Second

	for i := 0; i < maxRetries; i++ {
		resp, err := c.httpClient.Get(videoURL)
		if err != nil {
			lastErr = fmt.Errorf("failed to download video: %w", err)
			log.Printf("Attempt %d to download video for %s failed: %v. Retrying in %v...", i+1, shortcode, err, retryDelay)
			time.Sleep(retryDelay)
			retryDelay *= 2
			continue
		}

		if resp.StatusCode == http.StatusOK {
			defer resp.Body.Close()
			_, err = io.Copy(out, resp.Body)
			if err != nil {
				// If copy fails, we can't really recover, so just fail.
				return "", fmt.Errorf("failed to save video: %w", err)
			}
			return filePath, nil
		}

		resp.Body.Close() // close body if not OK
		lastErr = fmt.Errorf("unexpected status code while downloading: %d", resp.StatusCode)

		if resp.StatusCode == http.StatusTooManyRequests {
			log.Printf("Attempt %d: Rate limit hit downloading video for %s. Retrying in %v...", i+1, shortcode, retryDelay)
			time.Sleep(retryDelay)
			retryDelay *= 2
			continue
		}

		// Fail on other non-OK statuses
		return "", lastErr
	}

	return "", fmt.Errorf("failed to download video for %s after %d retries: %w", shortcode, maxRetries, lastErr)
}
