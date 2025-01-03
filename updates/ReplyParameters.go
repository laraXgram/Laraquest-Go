package updates

type ReplyParameters struct {
    Message_id int64 `json:"message_id"`
    Chat_id *IntOrString `json:"chat_id"`
    Allow_sending_without_reply bool `json:"allow_sending_without_reply"`
    Quote string `json:"quote"`
    Quote_parse_mode string `json:"quote_parse_mode"`
    Quote_entities *[]MessageEntity `json:"quote_entities"`
    Quote_position int64 `json:"quote_position"`
}