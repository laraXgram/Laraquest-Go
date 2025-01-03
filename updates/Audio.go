package updates

type Audio struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    Duration int64 `json:"duration"`
    Performer string `json:"performer"`
    Title string `json:"title"`
    File_name string `json:"file_name"`
    Mime_type string `json:"mime_type"`
    File_size int64 `json:"file_size"`
    Thumbnail *PhotoSize `json:"thumbnail"`
}