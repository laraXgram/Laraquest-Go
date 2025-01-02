package params

type SendPollParams struct { 
    Chat_id any `json:"chat_id"`
    Question any `json:"question"`
    Options any `json:"options"`
    Is_anonymous *any `json:"is_anonymous"`
    Type *any `json:"type"`
    Allows_multiple_answers *any `json:"allows_multiple_answers"`
    Question_parse_mode *any `json:"question_parse_mode"`
    Question_entities *any `json:"question_entities"`
    Correct_option_id *any `json:"correct_option_id"`
    Explanation *any `json:"explanation"`
    Explanation_parse_mode *any `json:"explanation_parse_mode"`
    Explanation_entities *any `json:"explanation_entities"`
    Open_period *any `json:"open_period"`
    Close_date *any `json:"close_date"`
    Is_closed *any `json:"is_closed"`
    Message_thread_id *any `json:"message_thread_id"`
    Reply_parameters *any `json:"reply_parameters"`
    Reply_markup *any `json:"reply_markup"`
    Protect_content *any `json:"protect_content"`
    Disable_notification *any `json:"disable_notification"`
    Business_connection_id *any `json:"business_connection_id"`
    Message_effect_id *any `json:"message_effect_id"`
    Allow_paid_broadcast *any `json:"allow_paid_broadcast"`
}