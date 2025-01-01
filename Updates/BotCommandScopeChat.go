package updates

type BotCommandScopeChat struct {
    Type string `json:"type"`
    Chat_id *IntOrString `json:"chat_id"`
}