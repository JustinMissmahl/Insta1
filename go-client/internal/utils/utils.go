package utils

import "regexp"

func ExtractShortcode(instagramURL string) string {
	postRegex := regexp.MustCompile(`/(p|reel|reels)/([a-zA-Z0-9_-]+)/?`)
	if matches := postRegex.FindStringSubmatch(instagramURL); len(matches) >= 3 && matches[2] != "" {
		return matches[2]
	}
	return ""
}
