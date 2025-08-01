package types

type IGGraphQLResponseDto struct {
	Data       DataDto       `json:"data"`
	Extensions ExtensionsDto `json:"extensions"`
	Status     string        `json:"status"`
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

type XdtShortcodeMediaDto struct {
	Typename                    string                           `json:"__typename"`
	IsXDTGraphMediaInterface    string                           `json:"__isXDTGraphMediaInterface"`
	ID                          string                           `json:"id"`
	Shortcode                   string                           `json:"shortcode"`
	ThumbnailSrc                string                           `json:"thumbnail_src"`
	Dimensions                  DimensionsDto                    `json:"dimensions"`
	GatingInfo                  interface{}                      `json:"gating_info"`
	FactCheckOverallRating      interface{}                      `json:"fact_check_overall_rating"`
	FactCheckInformation        interface{}                      `json:"fact_check_information"`
	SensitivityFrictionInfo     interface{}                      `json:"sensitivity_friction_info"`
	SharingFrictionInfo         SharingFrictionInfoDto           `json:"sharing_friction_info"`
	MediaOverlayInfo            interface{}                      `json:"media_overlay_info"`
	MediaPreview                string                           `json:"media_preview"`
	DisplayUrl                  string                           `json:"display_url"`
	DisplayResources            []DisplayResourceDto             `json:"display_resources"`
	AccessibilityCaption        interface{}                      `json:"accessibility_caption"`
	DashInfo                    DashInfoDto                      `json:"dash_info"`
	HasAudio                    bool                             `json:"has_audio"`
	VideoUrl                    string                           `json:"video_url"`
	VideoViewCount              int64                            `json:"video_view_count"`
	VideoPlayCount              int64                            `json:"video_play_count"`
	EncodingStatus              interface{}                      `json:"encoding_status"`
	IsPublished                 bool                             `json:"is_published"`
	ProductType                 string                           `json:"product_type"`
	Title                       string                           `json:"title"`
	VideoDuration               float64                          `json:"video_duration"`
	ClipsMusicAttributionInfo   ClipsMusicAttributionInfoDto     `json:"clips_music_attribution_info"`
	IsVideo                     bool                             `json:"is_video"`
	TrackingToken               string                           `json:"tracking_token"`
	UpcomingEvent               interface{}                      `json:"upcoming_event"`
	EdgeMediaToTaggedUser       EdgeMediaToCaptionClassDto       `json:"edge_media_to_tagged_user"`
	Owner                       XdtShortcodeMediaOwnerDto        `json:"owner"`
	EdgeMediaToCaption          EdgeMediaToCaptionClassDto       `json:"edge_media_to_caption"`
	CanSeeInsightsAsBrand       bool                             `json:"can_see_insights_as_brand"`
	CaptionIsEdited             bool                             `json:"caption_is_edited"`
	HasRankedComments           bool                             `json:"has_ranked_comments"`
	LikeAndViewCountsDisabled   bool                             `json:"like_and_view_counts_disabled"`
	EdgeMediaToParentComment    EdgeMediaToParentCommentClassDto `json:"edge_media_to_parent_comment"`
	EdgeMediaToHoistedComment   EdgeMediaToCaptionClassDto       `json:"edge_media_to_hoisted_comment"`
	EdgeMediaPreviewComment     EdgeMediaPreviewDto              `json:"edge_media_preview_comment"`
	CommentsDisabled            bool                             `json:"comments_disabled"`
	CommentingDisabledForViewer bool                             `json:"commenting_disabled_for_viewer"`
	TakenAtTimestamp            int64                            `json:"taken_at_timestamp"`
	EdgeMediaPreviewLike        EdgeMediaPreviewDto              `json:"edge_media_preview_like"`
	EdgeMediaToSponsorUser      EdgeMediaToCaptionClassDto       `json:"edge_media_to_sponsor_user"`
	IsAffiliate                 bool                             `json:"is_affiliate"`
	IsPaidPartnership           bool                             `json:"is_paid_partnership"`
	Location                    interface{}                      `json:"location"`
	NftAssetInfo                interface{}                      `json:"nft_asset_info"`
	ViewerHasLiked              bool                             `json:"viewer_has_liked"`
	ViewerHasSaved              bool                             `json:"viewer_has_saved"`
	ViewerHasSavedToCollection  bool                             `json:"viewer_has_saved_to_collection"`
	ViewerInPhotoOfYou          bool                             `json:"viewer_in_photo_of_you"`
	ViewerCanReshare            bool                             `json:"viewer_can_reshare"`
	IsAd                        bool                             `json:"is_ad"`
	EdgeWebMediaToRelatedMedia  EdgeMediaToCaptionClassDto       `json:"edge_web_media_to_related_media"`
	CoauthorProducers           []interface{}                    `json:"coauthor_producers"`
	PinnedForUsers              []interface{}                    `json:"pinned_for_users"`
}

type ClipsMusicAttributionInfoDto struct {
	ArtistName            string `json:"artist_name"`
	SongName              string `json:"song_name"`
	UsesOriginalAudio     bool   `json:"uses_original_audio"`
	ShouldMuteAudio       bool   `json:"should_mute_audio"`
	ShouldMuteAudioReason string `json:"should_mute_audio_reason"`
	AudioId               string `json:"audio_id"`
}

type DashInfoDto struct {
	IsDashEligible    bool   `json:"is_dash_eligible"`
	VideoDashManifest string `json:"video_dash_manifest"`
	NumberOfQualities int    `json:"number_of_qualities"`
}

type DimensionsDto struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type DisplayResourceDto struct {
	Src          string `json:"src"`
	ConfigWidth  int    `json:"config_width"`
	ConfigHeight int    `json:"config_height"`
}

type EdgeMediaPreviewDto struct {
	Count int                              `json:"count"`
	Edges []EdgeMediaPreviewCommentEdgeDto `json:"edges"`
}

type EdgeMediaToParentCommentClassDto struct {
	Count    int                              `json:"count"`
	PageInfo PageInfoDto                      `json:"page_info"`
	Edges    []EdgeMediaPreviewCommentEdgeDto `json:"edges"`
}

type PurpleNodeDto struct {
	ID                   string                            `json:"id"`
	Text                 string                            `json:"text"`
	CreatedAt            int64                             `json:"created_at"`
	DidReportAsSpam      bool                              `json:"did_report_as_spam"`
	Owner                NodeOwnerDto                      `json:"owner"`
	ViewerHasLiked       bool                              `json:"viewer_has_liked"`
	EdgeLikedBy          EdgeFollowedByClassDto            `json:"edge_liked_by"`
	IsRestrictedPending  bool                              `json:"is_restricted_pending"`
	EdgeThreadedComments *EdgeMediaToParentCommentClassDto `json:"edge_threaded_comments,omitempty"`
}

type EdgeMediaPreviewCommentEdgeDto struct {
	Node PurpleNodeDto `json:"node"`
}

type PageInfoDto struct {
	HasNextPage bool    `json:"has_next_page"`
	EndCursor   *string `json:"end_cursor"`
}

type EdgeFollowedByClassDto struct {
	Count int `json:"count"`
}

type NodeOwnerDto struct {
	ID            string `json:"id"`
	IsVerified    bool   `json:"is_verified"`
	ProfilePicUrl string `json:"profile_pic_url"`
	Username      string `json:"username"`
}

type EdgeMediaToCaptionClassDto struct {
	Edges []EdgeMediaToCaptionEdgeDto `json:"edges"`
}

type EdgeMediaToCaptionEdgeDto struct {
	Node FluffyNodeDto `json:"node"`
}

type FluffyNodeDto struct {
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
	ID        string `json:"id"`
}

type XdtShortcodeMediaOwnerDto struct {
	ID                        string                 `json:"id"`
	Username                  string                 `json:"username"`
	IsVerified                bool                   `json:"is_verified"`
	ProfilePicUrl             string                 `json:"profile_pic_url"`
	BlockedByViewer           bool                   `json:"blocked_by_viewer"`
	RestrictedByViewer        interface{}            `json:"restricted_by_viewer"`
	FollowedByViewer          bool                   `json:"followed_by_viewer"`
	FullName                  string                 `json:"full_name"`
	HasBlockedViewer          bool                   `json:"has_blocked_viewer"`
	IsEmbedsDisabled          bool                   `json:"is_embeds_disabled"`
	IsPrivate                 bool                   `json:"is_private"`
	IsUnpublished             bool                   `json:"is_unpublished"`
	RequestedByViewer         bool                   `json:"requested_by_viewer"`
	PassTieringRecommendation bool                   `json:"pass_tiering_recommendation"`
	EdgeOwnerToTimelineMedia  EdgeFollowedByClassDto `json:"edge_owner_to_timeline_media"`
	EdgeFollowedBy            EdgeFollowedByClassDto `json:"edge_followed_by"`
}

type SharingFrictionInfoDto struct {
	ShouldHaveSharingFriction bool        `json:"should_have_sharing_friction"`
	BloksAppUrl               interface{} `json:"bloks_app_url"`
}

type ExtensionsDto struct {
	IsFinal bool `json:"is_final"`
}
