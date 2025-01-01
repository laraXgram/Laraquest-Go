package updates

type MessageOriginChannel struct {
    Type string `json:"type"`
    Date int64 `json:"date"`
    Chat *Chat `json:"chat"`
    Message_id int64 `json:"message_id"`
    Author_signature string `json:"author_signature"`
}