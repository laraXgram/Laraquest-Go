package updates

type ChatMemberUpdated struct {
    Chat *Chat `json:"chat"`
    From *User `json:"from"`
    Date int64 `json:"date"`
    Old_chat_member *ChatMember `json:"old_chat_member"`
    New_chat_member *ChatMember `json:"new_chat_member"`
    Invite_link *ChatInviteLink `json:"invite_link"`
    Via_join_request bool `json:"via_join_request"`
    Via_chat_folder_invite_link bool `json:"via_chat_folder_invite_link"`
}