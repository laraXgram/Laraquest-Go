package updates

type InputMediaAnimation struct {
    Type string `json:"type"`
    Media string `json:"media"`
    Thumbnail *InputFileOrString `json:"thumbnail"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Show_caption_above_media bool `json:"show_caption_above_media"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    Duration int64 `json:"duration"`
    Has_spoiler bool `json:"has_spoiler"`
}