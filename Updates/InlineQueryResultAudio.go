package updates

type InlineQueryResultAudio struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Audio_url string `json:"audio_url"`
    Title string `json:"title"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Performer string `json:"performer"`
    Audio_duration int64 `json:"audio_duration"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}