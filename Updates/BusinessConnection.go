package updates

type BusinessConnection struct {
    Id string `json:"id"`
    User *User `json:"user"`
    User_chat_id int64 `json:"user_chat_id"`
    Date int64 `json:"date"`
    Can_reply bool `json:"can_reply"`
    Is_enabled bool `json:"is_enabled"`
}