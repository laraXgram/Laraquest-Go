package updates

type ChatJoinRequest struct {
    Chat *Chat `json:"chat"`
    From *User `json:"from"`
    User_chat_id int64 `json:"user_chat_id"`
    Date int64 `json:"date"`
    Bio string `json:"bio"`
    Invite_link *ChatInviteLink `json:"invite_link"`
}