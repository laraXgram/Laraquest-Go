package updates

type KeyboardButtonRequestChat struct {
    Request_id int64 `json:"request_id"`
    Chat_is_channel bool `json:"chat_is_channel"`
    Chat_is_forum bool `json:"chat_is_forum"`
    Chat_has_username bool `json:"chat_has_username"`
    Chat_is_created bool `json:"chat_is_created"`
    User_administrator_rights *ChatAdministratorRights `json:"user_administrator_rights"`
    Bot_administrator_rights *ChatAdministratorRights `json:"bot_administrator_rights"`
    Bot_is_member bool `json:"bot_is_member"`
    Request_title bool `json:"request_title"`
    Request_username bool `json:"request_username"`
    Request_photo bool `json:"request_photo"`
}