package params

type GetGameHighScoresParams struct { 
    User_id any `json:"user_id"`
    Chat_id *any `json:"chat_id"`
    Message_id *any `json:"message_id"`
    Inline_message_id *any `json:"inline_message_id"`
}