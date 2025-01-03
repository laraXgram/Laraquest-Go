package updates

type InlineQueryResultGame struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Game_short_name string `json:"game_short_name"`
    Reply_markup *InlineKeyboardMarkup `json:"reply_markup"`
}