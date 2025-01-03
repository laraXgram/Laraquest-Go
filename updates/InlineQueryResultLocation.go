package updates

type InlineQueryResultLocation struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Title string `json:"title"`
    Horizontal_accuracy float64 `json:"horizontal_accuracy"`
    Live_period int64 `json:"live_period"`
    Heading int64 `json:"heading"`
    Proximity_alert_radius int64 `json:"proximity_alert_radius"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
    Thumbnail_url string `json:"thumbnail_url"`
    Thumbnail_width int64 `json:"thumbnail_width"`
    Thumbnail_height int64 `json:"thumbnail_height"`
}