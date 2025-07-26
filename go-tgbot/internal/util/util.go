package util

import (
	"regexp"
)

var (
	instagramPostRegex = regexp.MustCompile(`(?:https?:\/\/)?(?:www\.)?instagram\.com\/(?:p|reel|reels)\/([a-zA-Z0-9_-]+)`)
)

func ExtractShortcode(url string) string {
	matches := instagramPostRegex.FindStringSubmatch(url)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}
