package utils

import (
	"regexp"
	"strings"
)

// IsShortcodePresent checks if a URL contains a valid Instagram shortcode
func IsShortcodePresent(url string) bool {
	regex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	matches := regex.FindStringSubmatch(url)
	return len(matches) >= 3 && matches[2] != ""
}

// GetPostShortcode extracts the shortcode from an Instagram URL
func GetPostShortcode(url string) string {
	regex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	matches := regex.FindStringSubmatch(url)

	if len(matches) >= 3 && matches[2] != "" {
		return matches[2]
	}

	return ""
}

// IsValidInstagramURL validates if a URL is a valid Instagram post/reel URL
func IsValidInstagramURL(url string) bool {
	if !strings.HasPrefix(url, "https://www.instagram.com") {
		return false
	}

	return IsShortcodePresent(url)
}
