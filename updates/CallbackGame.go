package updates

type CallbackGame struct {
    User_id int64 `json:"user_id"`
    Score int64 `json:"score"`
    Force bool `json:"force"`
    Disable_edit_message bool `json:"disable_edit_message"`
    Chat_id int64 `json:"chat_id"`
    Message_id int64 `json:"message_id"`
    Inline_message_id string `json:"inline_message_id"`
}