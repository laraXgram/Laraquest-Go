package updates

type InlineQueryResultCachedAudio struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Audio_file_id string `json:"audio_file_id"`
    Caption string `json:"caption"`
    Parse_mode string `json:"parse_mode"`
    Caption_entities *[]MessageEntity `json:"caption_entities"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}