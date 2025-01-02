package params

type UnpinChatMessageParams struct { 
    Chat_id any `json:"chat_id"`
    Message_id any `json:"message_id"`
    Business_connection_id *any `json:"business_connection_id"`
}