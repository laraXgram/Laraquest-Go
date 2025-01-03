package updates

type InlineQueryResultCachedDocument struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Title string `json:"title"`
    Document_file_id string `json:"document_file_id"`
    Description string `json:"description"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}