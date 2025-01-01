package updates

type File struct {
    File_id string `json:"file_id"`
    File_unique_id string `json:"file_unique_id"`
    File_size int64 `json:"file_size"`
    File_path string `json:"file_path"`
}