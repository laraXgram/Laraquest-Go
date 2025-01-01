package updates

type PassportElementErrorFiles struct {
    Source string `json:"source"`
    Type string `json:"type"`
    File_hashes []string `json:"file_hashes"`
    Message string `json:"message"`
}