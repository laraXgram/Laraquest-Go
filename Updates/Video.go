package updates

type Video struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    Duration int64 `json:"duration"`
    Thumbnail *PhotoSize `json:"thumbnail"`
    File_name string `json:"file_name"`
    Mime_type string `json:"mime_type"`
    File_size int64 `json:"file_size"`
}