package updates

type InlineQueryResultVoice struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Voice_url string `json:"voice_url"`
    Title string `json:"title"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Voice_duration int64 `json:"voice_duration"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}