package params

type SetGameScoreParams struct { 
    User_id any `json:"user_id"`
    Score any `json:"score"`
    Chat_id *any `json:"chat_id"`
    Message_id *any `json:"message_id"`
    Inline_message_id *any `json:"inline_message_id"`
    Disable_edit_message *any `json:"disable_edit_message"`
    Force *any `json:"force"`
}