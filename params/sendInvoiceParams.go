package params

type SendInvoiceParams struct { 
    Chat_id any `json:"chat_id"`
    Title any `json:"title"`
    Description any `json:"description"`
    Payload any `json:"payload"`
    Provider_token any `json:"provider_token"`
    Currency any `json:"currency"`
    Price any `json:"price"`
    Message_thread_id *any `json:"message_thread_id"`
    Max_tip_amount *any `json:"max_tip_amount"`
    Suggested_tip_amounts *any `json:"suggested_tip_amounts"`
    Start_parameter *any `json:"start_parameter"`
    Provider_data *any `json:"provider_data"`
    Photo_url *any `json:"photo_url"`
    Photo_size *any `json:"photo_size"`
    Photo_width *any `json:"photo_width"`
    Photo_height *any `json:"photo_height"`
    Need_name *any `json:"need_name"`
    Need_phone_number *any `json:"need_phone_number"`
    Need_email *any `json:"need_email"`
    Need_shipping_address *any `json:"need_shipping_address"`
    Send_phone_number_to_provider *any `json:"send_phone_number_to_provider"`
    Send_email_to_provider *any `json:"send_email_to_provider"`
    Is_flexible *any `json:"is_flexible"`
    Disable_notification *any `json:"disable_notification"`
    Protect_content *any `json:"protect_content"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}