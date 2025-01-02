package params

type SendPhotoParams struct { 
    Chat_id any `json:"chat_id"`
    Photo any `json:"photo"`
    Caption *any `json:"caption"`
    Parse_mode *any `json:"parse_mode"`
    Message_thread_id *any `json:"message_thread_id"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Has_spoiler *any `json:"has_spoiler"`
    Caption_entities *any `json:"caption_entities"`
    Show_caption_above_media *any `json:"show_caption_above_media"`
    Message_effect_id *any `json:"message_effect_id"`
    Business_connection_id *any `json:"business_connection_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}