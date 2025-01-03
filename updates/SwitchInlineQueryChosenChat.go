package updates

type SwitchInlineQueryChosenChat struct {
    Query string `json:"query"`
    Allow_user_chats bool `json:"allow_user_chats"`
    Allow_bot_chats bool `json:"allow_bot_chats"`
    Allow_group_chats bool `json:"allow_group_chats"`
    Allow_channel_chats bool `json:"allow_channel_chats"`
}