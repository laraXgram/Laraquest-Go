package params

type SetChatPermissionsParams struct { 
    Chat_id any `json:"chat_id"`
    Permissions any `json:"permissions"`
    Use_independent_chat_permissions *any `json:"use_independent_chat_permissions"`
}