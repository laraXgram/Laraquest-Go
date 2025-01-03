package updates

type InlineQueryResultCachedSticker struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Sticker_file_id string `json:"sticker_file_id"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
}