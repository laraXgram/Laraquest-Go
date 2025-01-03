package updates

type InlineQueryResultContact struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Phone_number string `json:"phone_number"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    Vcard string `json:"vcard"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
    Thumbnail_url string `json:"thumbnail_url"`
    Thumbnail_width int64 `json:"thumbnail_width"`
    Thumbnail_height int64 `json:"thumbnail_height"`
}