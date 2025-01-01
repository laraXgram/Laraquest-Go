package updates

type ChatMemberBanned struct {
    Status string `json:"status"`
    User *User `json:"user"`
    Until_date int64 `json:"until_date"`
}