package updates

type TextQuote struct {
    Text string `json:"text"`
    Entities *[]MessageEntity `json:"entities"`
    Position int64 `json:"position"`
    Is_manual bool `json:"is_manual"`
}