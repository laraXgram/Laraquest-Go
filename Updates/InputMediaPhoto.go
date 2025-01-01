package updates

type InputMediaPhoto struct {
    Type string `json:"type"`
    Media string `json:"media"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Show_caption_above_media bool `json:"show_caption_above_media"`
    Has_spoiler bool `json:"has_spoiler"`
}