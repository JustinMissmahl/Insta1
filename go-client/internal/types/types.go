package types

// InstagramResponse represents the API response structure
type InstagramResponse struct {
	Data struct {
		XdtShortcodeMedia struct {
			ID       string `json:"id"`
			IsVideo  bool   `json:"is_video"`
			VideoURL string `json:"video_url"`
			Owner    struct {
				Username string `json:"username"`
				FullName string `json:"full_name"`
			} `json:"owner"`
			VideoDuration float64 `json:"video_duration"`
		} `json:"xdt_shortcode_media"`
	} `json:"data"`
}
