package updates

type ChatShared struct {
    Request_id int64 `json:"request_id"`
    Chat_id int64 `json:"chat_id"`
    Title string `json:"title"`
    Username string `json:"username"`
    Photo *[]PhotoSize `json:"photo"`
}