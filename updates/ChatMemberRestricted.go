package updates

type ChatMemberRestricted struct {
    Status string `json:"status"`
    User *User `json:"user"`
    Is_member bool `json:"is_member"`
    Can_send_messages bool `json:"can_send_messages"`
    Can_send_audios bool `json:"can_send_audios"`
    Can_send_documents bool `json:"can_send_documents"`
    Can_send_photos bool `json:"can_send_photos"`
    Can_send_videos bool `json:"can_send_videos"`
    Can_send_video_notes bool `json:"can_send_video_notes"`
    Can_send_voice_notes bool `json:"can_send_voice_notes"`
    Can_send_polls bool `json:"can_send_polls"`
    Can_send_other_messages bool `json:"can_send_other_messages"`
    Can_add_web_page_previews bool `json:"can_add_web_page_previews"`
    Can_change_info bool `json:"can_change_info"`
    Can_invite_users bool `json:"can_invite_users"`
    Can_pin_messages bool `json:"can_pin_messages"`
    Can_manage_topics bool `json:"can_manage_topics"`
    Until_date int64 `json:"until_date"`
}