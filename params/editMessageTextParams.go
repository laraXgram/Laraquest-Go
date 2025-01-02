package params

type EditMessageTextParams struct { 
    Text any `json:"text"`
    Chat_id *any `json:"chat_id"`
    Message_id *any `json:"message_id"`
    Inline_message_id *any `json:"inline_message_id"`
    Parse_mode *any `json:"parse_mode"`
    Entities *any `json:"entities"`
    Link_preview_options *any `json:"link_preview_options"`
    Reply_markup *any `json:"reply_markup"`
    Business_connection_id *any `json:"business_connection_id"`
}