package params

type SetUserEmojiStatusParams struct { 
    User_id any `json:"user_id"`
    Emoji_status_custom_emoji_id *any `json:"emoji_status_custom_emoji_id"`
    Emoji_status_expiration_date *any `json:"emoji_status_expiration_date"`
}