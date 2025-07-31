package utils

import "regexp"

// extractShortcode extracts shortcode from Instagram URL
func ExtractShortcode(instagramURL string) string {
	regex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	matches := regex.FindStringSubmatch(instagramURL)

	if len(matches) >= 3 && matches[2] != "" {
		return matches[2]
	}

	return ""
}
