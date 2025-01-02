package params

type SendLocationParams struct { 
    Chat_id any `json:"chat_id"`
    Latitude any `json:"latitude"`
    Longitude any `json:"longitude"`
    Horizontal_accuracy *any `json:"horizontal_accuracy"`
    Live_period *any `json:"live_period"`
    Heading *any `json:"heading"`
    Proximity_alert_radius *any `json:"proximity_alert_radius"`
    Message_thread_id *any `json:"message_thread_id"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}