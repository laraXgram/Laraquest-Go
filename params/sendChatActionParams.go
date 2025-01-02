package params

type SendChatActionParams struct { 
    Chat_id any `json:"chat_id"`
    Action any `json:"action"`
    Message_thread_id *any `json:"message_thread_id"`
    Business_connection_id *any `json:"business_connection_id"`
}