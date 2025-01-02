package params

type ForwardMessagesParams struct { 
    Chat_id any `json:"chat_id"`
    From_chat_id any `json:"from_chat_id"`
    Message_ids any `json:"message_ids"`
    Message_thread_id *any `json:"message_thread_id"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
}