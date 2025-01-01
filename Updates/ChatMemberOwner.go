package updates

type ChatMemberOwner struct {
    Status string `json:"status"`
    User *User `json:"user"`
    Is_anonymous bool `json:"is_anonymous"`
    Custom_title string `json:"custom_title"`
}