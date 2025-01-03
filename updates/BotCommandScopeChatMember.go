package updates

type BotCommandScopeChatMember struct {
    Type string `json:"type"`
    Chat_id *IntOrString `json:"chat_id"`
    User_id int64 `json:"user_id"`
}