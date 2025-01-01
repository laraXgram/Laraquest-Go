package updates

type InlineQueryResultCachedGif struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Gif_file_id string `json:"gif_file_id"`
    Title string `json:"title"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Show_caption_above_media bool `json:"show_caption_above_media"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}