package updates

type BackgroundTypeWallpaper struct {
    Type string `json:"type"`
    Document *Document `json:"document"`
    Dark_theme_dimming int64 `json:"dark_theme_dimming"`
    Is_blurred bool `json:"is_blurred"`
    Is_moving bool `json:"is_moving"`
}