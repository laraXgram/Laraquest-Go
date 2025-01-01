package updates

type InputPaidMediaVideo struct {
    Type string `json:"type"`
    Media string `json:"media"`
    Thumbnail *InputFileOrString `json:"thumbnail"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    Duration int64 `json:"duration"`
    Supports_streaming bool `json:"supports_streaming"`
}