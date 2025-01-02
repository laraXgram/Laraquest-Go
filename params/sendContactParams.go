package params

type SendContactParams struct { 
    Chat_id any `json:"chat_id"`
    Phone_number any `json:"phone_number"`
    First_name any `json:"first_name"`
    Last_name *any `json:"last_name"`
    Vcard *any `json:"vcard"`
    Message_thread_id *any `json:"message_thread_id"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}