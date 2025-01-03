package updates

type InlineQueryResultDocument struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Title string `json:"title"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Document_url string `json:"document_url"`
    Mime_type string `json:"mime_type"`
    Description string `json:"description"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
    Thumbnail_url string `json:"thumbnail_url"`
    Thumbnail_width int64 `json:"thumbnail_width"`
    Thumbnail_height int64 `json:"thumbnail_height"`
}