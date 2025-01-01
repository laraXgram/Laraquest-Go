package updates

type InlineQueryResultsButton struct {
    Text string `json:"text"`
    Web_app *WebAppInfo `json:"web_app"`
    Start_parameter string `json:"start_parameter"`
}