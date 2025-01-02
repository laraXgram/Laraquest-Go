package params

type SendGiftParams struct { 
    User_id any `json:"user_id"`
    Gift_id any `json:"gift_id"`
    Pay_for_upgrade *any `json:"pay_for_upgrade"`
    Text *any `json:"text"`
    Text_parse_mode *any `json:"text_parse_mode"`
    Text_entities *any `json:"text_entities"`
}