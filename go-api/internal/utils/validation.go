package utils

import (
	"log"
	"regexp"
	"strings"
)

// IsShortcodePresent checks if a URL contains a valid Instagram shortcode
func IsShortcodePresent(url string) bool {
	// Check for regular posts and reels
	postRegex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	if matches := postRegex.FindStringSubmatch(url); len(matches) >= 3 && matches[2] != "" {
		log.Printf("DEBUG: Found post/reel URL pattern: %v", matches)
		return true
	}
	log.Printf("DEBUG: No valid URL pattern found in: %s", url)
	return false
}

// GetPostShortcode extracts the shortcode from an Instagram URL
func GetPostShortcode(url string) string {
	// Check for regular posts and reels
	postRegex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	if matches := postRegex.FindStringSubmatch(url); len(matches) >= 3 && matches[2] != "" {
		log.Printf("DEBUG: Extracted post/reel shortcode: %s", matches[2])
		return matches[2]
	}
	log.Printf("DEBUG: Failed to extract shortcode from URL: %s", url)
	return ""
}

// IsValidInstagramURL validates if a URL is a valid Instagram post/reel URL
func IsValidInstagramURL(url string) bool {
	if !strings.HasPrefix(url, "https://www.instagram.com") {
		return false
	}
	return IsShortcodePresent(url)
}
