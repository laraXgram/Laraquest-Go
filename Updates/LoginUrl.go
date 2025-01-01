package updates

type LoginUrl struct {
    Url string `json:"url"`
    Forward_text string `json:"forward_text"`
    Bot_username string `json:"bot_username"`
    Request_write_access bool `json:"request_write_access"`
}