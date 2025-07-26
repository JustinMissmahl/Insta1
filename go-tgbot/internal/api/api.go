package api

import (
	"encoding/json"
	"fmt"
	"go-tgbot/internal/types"
	"io"
	"net/http"
	"os"
	"path/filepath"
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

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

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

	resp, err := c.httpClient.Get(videoURL)
	if err != nil {
		return "", fmt.Errorf("failed to download video: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code while downloading: %d", resp.StatusCode)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save video: %w", err)
	}

	return filePath, nil
}
