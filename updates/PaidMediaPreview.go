package updates

type PaidMediaPreview struct {
    Type string `json:"type"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    Duration int64 `json:"duration"`
}