package updates

type VideoNote struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    Length int64 `json:"length"`
    Duration int64 `json:"duration"`
    Thumbnail *PhotoSize `json:"thumbnail"`
    File_size int64 `json:"file_size"`
}