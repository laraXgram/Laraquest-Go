package params

type UnbanChatMemberParams struct { 
    Chat_id any `json:"chat_id"`
    User_id any `json:"user_id"`
    Only_if_banned *any `json:"only_if_banned"`
}