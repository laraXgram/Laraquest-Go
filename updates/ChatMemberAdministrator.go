package updates

type ChatMemberAdministrator struct {
    Status string `json:"status"`
    User *User `json:"user"`
    Can_be_edited bool `json:"can_be_edited"`
    Is_anonymous bool `json:"is_anonymous"`
    Can_manage_chat bool `json:"can_manage_chat"`
    Can_delete_messages bool `json:"can_delete_messages"`
    Can_manage_video_chats bool `json:"can_manage_video_chats"`
    Can_restrict_members bool `json:"can_restrict_members"`
    Can_promote_members bool `json:"can_promote_members"`
    Can_change_info bool `json:"can_change_info"`
    Can_invite_users bool `json:"can_invite_users"`
    Can_post_stories bool `json:"can_post_stories"`
    Can_edit_stories bool `json:"can_edit_stories"`
    Can_delete_stories bool `json:"can_delete_stories"`
    Can_post_messages bool `json:"can_post_messages"`
    Can_edit_messages bool `json:"can_edit_messages"`
    Can_pin_messages bool `json:"can_pin_messages"`
    Can_manage_topics bool `json:"can_manage_topics"`
    Custom_title string `json:"custom_title"`
}