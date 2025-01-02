package params

type BanChatMemberParams struct { 
    Chat_id any `json:"chat_id"`
    User_id any `json:"user_id"`
    Until_date *any `json:"until_date"`
    Revoke_messages *any `json:"revoke_messages"`
}