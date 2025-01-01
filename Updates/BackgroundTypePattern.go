package updates

type BackgroundTypePattern struct {
    Type string `json:"type"`
    Document *Document `json:"document"`
    Fill *BackgroundFill `json:"fill"`
    Intensity int64 `json:"intensity"`
    Is_inverted bool `json:"is_inverted"`
    Is_moving bool `json:"is_moving"`
}