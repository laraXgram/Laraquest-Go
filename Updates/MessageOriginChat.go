package updates

type MessageOriginChat struct {
    Type string `json:"type"`
    Date int64 `json:"date"`
    Sender_chat *Chat `json:"sender_chat"`
    Author_signature string `json:"author_signature"`
}