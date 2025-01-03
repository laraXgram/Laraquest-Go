package updates

type PhotoSize struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    File_size int64 `json:"file_size"`
}