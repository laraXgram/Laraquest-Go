package params

type PinChatMessageParams struct { 
    Chat_id any `json:"chat_id"`
    Message_id any `json:"message_id"`
    Disable_notification *any `json:"disable_notification"`
    Business_connection_id *any `json:"business_connection_id"`
}