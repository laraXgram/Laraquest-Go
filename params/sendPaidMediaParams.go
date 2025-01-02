package params

type SendPaidMediaParams struct { 
    Chat_id any `json:"chat_id"`
    Star_count any `json:"star_count"`
    Media any `json:"media"`
    Payload *any `json:"payload"`
    Caption *any `json:"caption"`
    Pars_mode *any `json:"pars_mode"`
    Caption_entities *any `json:"caption_entities"`
    Show_caption_above_media *any `json:"show_caption_above_media"`
    Disable_notification *any `json:"disable_notification"`
    Protect_content *any `json:"protect_content"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Business_connection_id *any `json:"business_connection_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}