package params

type CopyMessageParams struct { 
    Chat_id any `json:"chat_id"`
    From_chat_id any `json:"from_chat_id"`
    Message_id any `json:"message_id"`
    Parse_mode *any `json:"parse_mode"`
    Message_thread_id *any `json:"message_thread_id"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Caption *any `json:"caption"`
    Caption_entities *any `json:"caption_entities"`
    Show_caption_above_media *any `json:"show_caption_above_media"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}