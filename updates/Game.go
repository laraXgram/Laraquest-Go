package updates

type Game struct {
    Title string `json:"title"`
    Description string `json:"description"`
    Photo *[]PhotoSize `json:"photo"`
    Text string `json:"text"`
    Text_entities *[]MessageEntity `json:"text_entities"`
    Animation *Animation `json:"animation"`
}