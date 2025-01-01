package updates

type PassportElementErrorDataField struct {
    Source string `json:"source"`
    Type string `json:"type"`
    Field_name string `json:"field_name"`
    Data_hash string `json:"data_hash"`
    Message string `json:"message"`
}