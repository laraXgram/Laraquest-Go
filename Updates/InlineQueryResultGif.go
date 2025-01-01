package updates

type InlineQueryResultGif struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Gif_url string `json:"gif_url"`
    Gif_width int64 `json:"gif_width"`
    Gif_height int64 `json:"gif_height"`
    Gif_duration int64 `json:"gif_duration"`
    Thumbnail_url string `json:"thumbnail_url"`
    Thumbnail_mime_type string `json:"thumbnail_mime_type"`
    Title string `json:"title"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Show_caption_above_media bool `json:"show_caption_above_media"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}