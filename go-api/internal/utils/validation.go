package utils

import (
	"log"
	"regexp"
	"strings"
)

func IsShortcodePresent(url string) bool {
	postRegex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	if matches := postRegex.FindStringSubmatch(url); len(matches) >= 3 && matches[2] != "" {
		log.Printf("DEBUG: Found post/reel URL pattern: %v", matches)
		return true
	}
	log.Printf("DEBUG: No valid URL pattern found in: %s", url)
	return false
}

func GetPostShortcode(url string) string {
	postRegex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	if matches := postRegex.FindStringSubmatch(url); len(matches) >= 3 && matches[2] != "" {
		log.Printf("DEBUG: Extracted post/reel shortcode: %s", matches[2])
		return matches[2]
	}
	log.Printf("DEBUG: Failed to extract shortcode from URL: %s", url)
	return ""
}

func IsValidInstagramURL(url string) bool {
	if !strings.HasPrefix(url, "https://www.instagram.com") {
		return false
	}
	return IsShortcodePresent(url)
}
