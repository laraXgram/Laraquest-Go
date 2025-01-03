package updates

type ChatMemberLeft struct {
    Status string `json:"status"`
    User *User `json:"user"`
}