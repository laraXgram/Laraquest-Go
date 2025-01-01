package updates

type ForumTopic struct {
    Message_thread_id int64 `json:"message_thread_id"`
    Name string `json:"name"`
    Icon_color int64 `json:"icon_color"`
    Icon_custom_emoji_id string `json:"icon_custom_emoji_id"`
}