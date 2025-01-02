package params

type CreateInvoiceLinkParams struct { 
    Title any `json:"title"`
    Description any `json:"description"`
    Payload any `json:"payload"`
    Provider_token any `json:"provider_token"`
    Currency any `json:"currency"`
    Price any `json:"price"`
    Subscription_period *any `json:"subscription_period"`
    Max_tip_amount *any `json:"max_tip_amount"`
    Suggested_tip_amounts *any `json:"suggested_tip_amounts"`
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
    Business_connection_id *any `json:"business_connection_id"`
}