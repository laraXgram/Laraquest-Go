package updates

type InlineQuery struct {
    Id string `json:"id"`
    From *User `json:"from"`
    Query string `json:"query"`
    Offset string `json:"offset"`
    Chat_type string `json:"chat_type"`
    Location *Location `json:"location"`
}