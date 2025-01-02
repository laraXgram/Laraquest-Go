package params

type CreateForumTopicParams struct { 
    Chat_id any `json:"chat_id"`
    Name any `json:"name"`
    Icon_color *any `json:"icon_color"`
    Icon_custom_emoji_id *any `json:"icon_custom_emoji_id"`
}