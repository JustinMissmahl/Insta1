package types

type IGGraphQLResponseDto struct {
	Data   DataDto `json:"data"`
	Status string  `json:"status"`
}

type DataDto struct {
	XdtShortcodeMedia XdtShortcodeMediaDto `json:"xdt_shortcode_media"`
	User              UserDto              `json:"user,omitempty"`
}

type UserDto struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	FullName      string `json:"full_name"`
	IsPrivate     bool   `json:"is_private"`
	IsVerified    bool   `json:"is_verified"`
	ProfilePicUrl string `json:"profile_pic_url"`
}

type DimensionsDto struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type XdtShortcodeMediaDto struct {
	ID               string                    `json:"id"`
	Shortcode        string                    `json:"shortcode"`
	IsVideo          bool                      `json:"is_video"`
	VideoURL         string                    `json:"video_url"`
	VideoDuration    float64                   `json:"video_duration"`
	Owner            XdtShortcodeMediaOwnerDto `json:"owner"`
	DisplayUrl       string                    `json:"display_url"`
	DisplayResources []DisplayResourceDto      `json:"display_resources"`
}

type DisplayResourceDto struct {
	Src          string `json:"src"`
	ConfigWidth  int    `json:"config_width"`
	ConfigHeight int    `json:"config_height"`
}

type XdtShortcodeMediaOwnerDto struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	FullName string `json:"full_name"`
}
