package params

type SendGameParams struct { 
    Chat_id any `json:"chat_id"`
    Game_short_name any `json:"game_short_name"`
    Message_thread_id *any `json:"message_thread_id"`
    Disable_notification *any `json:"disable_notification"`
    Protect_content *any `json:"protect_content"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}