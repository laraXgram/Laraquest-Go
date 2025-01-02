package params

type SendVideoNoteParams struct { 
    Chat_id any `json:"chat_id"`
    Video_note any `json:"video_note"`
    Message_thread_id *any `json:"message_thread_id"`
    Duration *any `json:"duration"`
    Length *any `json:"length"`
    Thumbnail *any `json:"thumbnail"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}