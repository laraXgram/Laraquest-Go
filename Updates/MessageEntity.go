package updates

type MessageEntity struct {
    Type string `json:"type"`
    Offset int64 `json:"offset"`
    Length int64 `json:"length"`
    Url string `json:"url"`
    User *User `json:"user"`
    Language string `json:"language"`
    Custom_emoji_id string `json:"custom_emoji_id"`
}