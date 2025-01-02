package params

type EditChatInviteLinkParams struct { 
    Chat_id any `json:"chat_id"`
    Invite_link any `json:"invite_link"`
    Name *any `json:"name"`
    Expire_date *any `json:"expire_date"`
    Member_limit *any `json:"member_limit"`
    Creates_join_request *any `json:"creates_join_request"`
}