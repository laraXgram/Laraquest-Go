package updates

type ForumTopicCreated struct {
    Name string `json:"name"`
    Icon_color int64 `json:"icon_color"`
    Icon_custom_emoji_id string `json:"icon_custom_emoji_id"`
}