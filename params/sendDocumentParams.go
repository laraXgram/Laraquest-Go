package params

type SendDocumentParams struct { 
    Chat_id any `json:"chat_id"`
    Document any `json:"document"`
    Caption *any `json:"caption"`
    Parse_mode *any `json:"parse_mode"`
    Message_thread_id *any `json:"message_thread_id"`
    Thumbnail *any `json:"thumbnail"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Caption_entities *any `json:"caption_entities"`
    Business_connection_id *any `json:"business_connection_id"`
    Disable_content_type_detection *any `json:"disable_content_type_detection"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}