package params

type EditForumTopicParams struct { 
    Chat_id any `json:"chat_id"`
    Message_thread_id any `json:"message_thread_id"`
    Name *any `json:"name"`
    Icon_custom_emoji_id *any `json:"icon_custom_emoji_id"`
}