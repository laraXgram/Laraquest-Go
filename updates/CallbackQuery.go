package updates

type CallbackQuery struct {
    Id string `json:"id"`
    From *User `json:"from"`
    Message *MaybeInaccessibleMessage `json:"message"`
    Inline_message_id string `json:"inline_message_id"`
    Chat_instance string `json:"chat_instance"`
    Data string `json:"data"`
    Game_short_name string `json:"game_short_name"`
}