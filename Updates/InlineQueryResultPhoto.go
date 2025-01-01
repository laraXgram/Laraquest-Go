package updates

type InlineQueryResultPhoto struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Photo_url string `json:"photo_url"`
    Thumbnail_url string `json:"thumbnail_url"`
    Photo_width int64 `json:"photo_width"`
    Photo_height int64 `json:"photo_height"`
    Title string `json:"title"`
    Description string `json:"description"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Show_caption_above_media bool `json:"show_caption_above_media"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}