package updates

type BackgroundTypeFill struct {
    Type string `json:"type"`
    Fill *BackgroundFill `json:"fill"`
    Dark_theme_dimming int64 `json:"dark_theme_dimming"`
}