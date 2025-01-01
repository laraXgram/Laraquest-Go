package updates

type ChosenInlineResult struct {
    Result_id string `json:"result_id"`
    From *User `json:"from"`
    Location *Location `json:"location"`
    Inline_message_id string `json:"inline_message_id"`
    Query string `json:"query"`
}