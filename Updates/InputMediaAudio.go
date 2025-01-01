package updates

type InputMediaAudio struct {
    Type string `json:"type"`
    Media string `json:"media"`
    Thumbnail *InputFileOrString `json:"thumbnail"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Duration int64 `json:"duration"`
    Performer string `json:"performer"`
    Title string `json:"title"`
}