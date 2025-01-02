package params

type EditMessageLiveLocationParams struct { 
    Latitude any `json:"latitude"`
    Longitude any `json:"longitude"`
    Chat_id *any `json:"chat_id"`
    Message_id *any `json:"message_id"`
    Inline_message_id *any `json:"inline_message_id"`
    Horizontal_accuracy *any `json:"horizontal_accuracy"`
    Heading *any `json:"heading"`
    Proximity_alert_radius *any `json:"proximity_alert_radius"`
    Reply_markup *any `json:"reply_markup"`
    Live_period *any `json:"live_period"`
    Business_connection_id *any `json:"business_connection_id"`
}