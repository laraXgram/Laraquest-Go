package updates

type InlineQueryResultVideo struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Video_url string `json:"video_url"`
    Mime_type string `json:"mime_type"`
    Thumbnail_url string `json:"thumbnail_url"`
    Title string `json:"title"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Show_caption_above_media bool `json:"show_caption_above_media"`
    Video_width int64 `json:"video_width"`
    Video_height int64 `json:"video_height"`
    Video_duration int64 `json:"video_duration"`
    Description string `json:"description"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}