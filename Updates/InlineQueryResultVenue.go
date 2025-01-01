package updates

type InlineQueryResultVenue struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Title string `json:"title"`
    Address string `json:"address"`
    Foursquare_id string `json:"foursquare_id"`
    Foursquare_type string `json:"foursquare_type"`
    Google_place_id string `json:"google_place_id"`
    Google_place_type string `json:"google_place_type"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
    Input_message_content *InputMessageContent `json:"input_message_content"`
    Thumbnail_url string `json:"thumbnail_url"`
    Thumbnail_width int64 `json:"thumbnail_width"`
    Thumbnail_height int64 `json:"thumbnail_height"`
}