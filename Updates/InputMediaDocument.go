package updates

type InputMediaDocument struct {
    Type string `json:"type"`
    Media string `json:"media"`
    Thumbnail *InputFileOrString `json:"thumbnail"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Disable_content_type_detection bool `json:"disable_content_type_detection"`
}