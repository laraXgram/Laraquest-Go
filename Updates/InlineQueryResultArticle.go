package updates

type InlineQueryResultArticle struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Title string `json:"title"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Url string `json:"url"`
    Description string `json:"description"`
    Thumbnail_url string `json:"thumbnail_url"`
    Thumbnail_width int64 `json:"thumbnail_width"`
    Thumbnail_height int64 `json:"thumbnail_height"`
}