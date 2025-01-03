package updates

type InputPollOption struct {
    Text string `json:"text"`
    Text_parse_mode string `json:"text_parse_mode"`
    Text_entities *[]MessageEntity `json:"text_entities"`
}