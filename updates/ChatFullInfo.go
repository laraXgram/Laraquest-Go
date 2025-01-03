package updates

type ChatFullInfo struct {
    Id int64 `json:"id"`
    Type string `json:"type"`
    Title string `json:"title"`
    Username string `json:"username"`
    First_name string `json:"first_name"`
    Last_name string `json:"last_name"`
    Is_forum bool `json:"is_forum"`
    Accent_color_id int64 `json:"accent_color_id"`
    Max_reaction_count int64 `json:"max_reaction_count"`
    Photo *ChatPhoto `json:"photo"`
    Active_usernames []string `json:"active_usernames"`
    Birthdate *Birthdate `json:"birthdate"`
    Business_intro *BusinessIntro `json:"business_intro"`
    Business_location *BusinessLocation `json:"business_location"`
    Business_opening_hours *BusinessOpeningHours `json:"business_opening_hours"`
    Personal_chat *Chat `json:"personal_chat"`
    Available_reactions *[]ReactionType `json:"available_reactions"`
    Background_custom_emoji_id string `json:"background_custom_emoji_id"`
    Profile_accent_color_id int64 `json:"profile_accent_color_id"`
    Profile_background_custom_emoji_id string `json:"profile_background_custom_emoji_id"`
    Emoji_status_custom_emoji_id string `json:"emoji_status_custom_emoji_id"`
    Emoji_status_expiration_date int64 `json:"emoji_status_expiration_date"`
    Bio string `json:"bio"`
    Has_private_forwards bool `json:"has_private_forwards"`
    Has_restricted_voice_and_video_messages bool `json:"has_restricted_voice_and_video_messages"`
    Join_to_send_messages bool `json:"join_to_send_messages"`
    Join_by_request bool `json:"join_by_request"`
    Description string `json:"description"`
    Invite_link string `json:"invite_link"`
    Pinned_message *Message `json:"pinned_message"`
    Permissions *ChatPermissions `json:"permissions"`
    Can_send_paid_media bool `json:"can_send_paid_media"`
    Slow_mode_delay int64 `json:"slow_mode_delay"`
    Unrestrict_boost_count int64 `json:"unrestrict_boost_count"`
    Message_auto_delete_time int64 `json:"message_auto_delete_time"`
    Has_aggressive_anti_spam_enabled bool `json:"has_aggressive_anti_spam_enabled"`
    Has_hidden_members bool `json:"has_hidden_members"`
    Has_protected_content bool `json:"has_protected_content"`
    Has_visible_history bool `json:"has_visible_history"`
    Sticker_set_name string `json:"sticker_set_name"`
    Can_set_sticker_set bool `json:"can_set_sticker_set"`
    Custom_emoji_sticker_set_name string `json:"custom_emoji_sticker_set_name"`
    Linked_chat_id int64 `json:"linked_chat_id"`
    Location *ChatLocation `json:"location"`
}