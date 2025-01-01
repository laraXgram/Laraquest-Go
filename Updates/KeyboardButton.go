package updates

type KeyboardButton struct {
    Text string `json:"text"`
    Request_users *KeyboardButtonRequestUsers `json:"request_users"`
    Request_chat *KeyboardButtonRequestChat `json:"request_chat"`
    Request_contact bool `json:"request_contact"`
    Request_location bool `json:"request_location"`
    Request_poll *KeyboardButtonPollType `json:"request_poll"`
    Web_app *WebAppInfo `json:"web_app"`
}