package params

type RestrictChatMemberParams struct { 
    Chat_id any `json:"chat_id"`
    User_id any `json:"user_id"`
    Permissions any `json:"permissions"`
    Until_date *any `json:"until_date"`
    Use_independent_chat_permissions *any `json:"use_independent_chat_permissions"`
}