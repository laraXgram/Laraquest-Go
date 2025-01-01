package updates

type MenuButtonWebApp struct {
    Type string `json:"type"`
    Text string `json:"text"`
    Web_app *WebAppInfo `json:"web_app"`
}