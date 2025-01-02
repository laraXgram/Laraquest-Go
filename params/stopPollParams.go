package params

type StopPollParams struct { 
    Chat_id any `json:"chat_id"`
    Message_id any `json:"message_id"`
    Reply_markup *any `json:"reply_markup"`
    Business_connection_id *any `json:"business_connection_id"`
}