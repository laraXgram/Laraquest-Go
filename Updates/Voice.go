package updates

type Voice struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    Duration int64 `json:"duration"`
    Mime_type string `json:"mime_type"`
    File_size int64 `json:"file_size"`
}