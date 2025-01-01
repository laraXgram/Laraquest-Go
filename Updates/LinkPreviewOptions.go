package updates

type LinkPreviewOptions struct {
    Is_disabled bool `json:"is_disabled"`
    Url string `json:"url"`
    Prefer_small_media bool `json:"prefer_small_media"`
    Prefer_large_media bool `json:"prefer_large_media"`
    Show_above_text bool `json:"show_above_text"`
}