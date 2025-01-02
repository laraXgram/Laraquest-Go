package params

type SendVenueParams struct { 
    Chat_id any `json:"chat_id"`
    Latitude any `json:"latitude"`
    Longitude any `json:"longitude"`
    Title any `json:"title"`
    Address any `json:"address"`
    Foursquare_id *any `json:"foursquare_id"`
    Foursquare_type *any `json:"foursquare_type"`
    Google_place_id *any `json:"google_place_id"`
    Google_place_type *any `json:"google_place_type"`
    Message_thread_id *any `json:"message_thread_id"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}