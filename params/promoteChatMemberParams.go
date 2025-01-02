package params

type PromoteChatMemberParams struct { 
    Chat_id any `json:"chat_id"`
    User_id any `json:"user_id"`
    Is_anonymous *any `json:"is_anonymous"`
    Can_manage_chat *any `json:"can_manage_chat"`
    Can_delete_messages *any `json:"can_delete_messages"`
    Can_manage_video_chats *any `json:"can_manage_video_chats"`
    Can_restrict_members *any `json:"can_restrict_members"`
    Can_promote_members *any `json:"can_promote_members"`
    Can_change_info *any `json:"can_change_info"`
    Can_invite_users *any `json:"can_invite_users"`
    Can_post_stories *any `json:"can_post_stories"`
    Can_edit_stories *any `json:"can_edit_stories"`
    Can_delete_stories *any `json:"can_delete_stories"`
    Can_post_messages *any `json:"can_post_messages"`
    Can_edit_messages *any `json:"can_edit_messages"`
    Can_pin_messages *any `json:"can_pin_messages"`
    Can_manage_topics *any `json:"can_manage_topics"`
}