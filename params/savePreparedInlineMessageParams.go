package params

type SavePreparedInlineMessageParams struct { 
    User_id any `json:"user_id"`
    Result any `json:"result"`
    Allow_user_chats *any `json:"allow_user_chats"`
    Allow_bot_chats *any `json:"allow_bot_chats"`
    Allow_group_chats *any `json:"allow_group_chats"`
    Allow_channel_chats *any `json:"allow_channel_chats"`
}