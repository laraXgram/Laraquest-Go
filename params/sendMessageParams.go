package params

type SendMessageParams struct { 
    Chat_id any `json:"chat_id"`
    Text any `json:"text"`
    Parse_mode *any `json:"parse_mode"`
    Message_thread_id *any `json:"message_thread_id"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Link_preview_options *any `json:"link_preview_options"`
    Entities *any `json:"entities"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}