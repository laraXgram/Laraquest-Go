package updates

type MessageOriginUser struct {
    Type string `json:"type"`
    Date int64 `json:"date"`
    Sender_user *User `json:"sender_user"`
}