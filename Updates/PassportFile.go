package updates

type PassportFile struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    File_size int64 `json:"file_size"`
    File_date int64 `json:"file_date"`
}