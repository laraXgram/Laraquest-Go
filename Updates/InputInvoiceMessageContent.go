package updates

type InputInvoiceMessageContent struct {
    Title string `json:"title"`
    Description string `json:"description"`
    Payload string `json:"payload"`
    Provider_token string `json:"provider_token"`
    Currency string `json:"currency"`
    Prices *[]LabeledPrice `json:"prices"`
    Max_tip_amount int64 `json:"max_tip_amount"`
    Suggested_tip_amounts []int64 `json:"suggested_tip_amounts"`
    Provider_data string `json:"provider_data"`
    Photo_url string `json:"photo_url"`
    Photo_size int64 `json:"photo_size"`
    Photo_width int64 `json:"photo_width"`
    Photo_height int64 `json:"photo_height"`
    Need_name bool `json:"need_name"`
    Need_phone_number bool `json:"need_phone_number"`
    Need_email bool `json:"need_email"`
    Need_shipping_address bool `json:"need_shipping_address"`
    Send_phone_number_to_provider bool `json:"send_phone_number_to_provider"`
    Send_email_to_provider bool `json:"send_email_to_provider"`
    Is_flexible bool `json:"is_flexible"`
}