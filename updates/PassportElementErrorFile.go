package updates

type PassportElementErrorFile struct {
    Source string `json:"source"`
    Type string `json:"type"`
    File_hash string `json:"file_hash"`
    Message string `json:"message"`
}