package updates

type InlineQueryResultCachedVoice struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Voice_file_id string `json:"voice_file_id"`
    Title string `json:"title"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}