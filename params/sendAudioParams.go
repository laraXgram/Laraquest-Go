package params

type SendAudioParams struct { 
    Chat_id any `json:"chat_id"`
    Audio any `json:"audio"`
    Caption *any `json:"caption"`
    Parse_mode *any `json:"parse_mode"`
    Message_thread_id *any `json:"message_thread_id"`
    Duration *any `json:"duration"`
    Performer *any `json:"performer"`
    Title *any `json:"title"`
    Thumbnail *any `json:"thumbnail"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Caption_entities *any `json:"caption_entities"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}