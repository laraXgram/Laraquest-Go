package updates

type InlineKeyboardButton struct {
    Text string `json:"text"`
    Url string `json:"url"`
    Callback_data string `json:"callback_data"`
    Web_app *WebAppInfo `json:"web_app"`
    Login_url *LoginUrl `json:"login_url"`
    Switch_inline_query string `json:"switch_inline_query"`
    Switch_inline_query_current_chat string `json:"switch_inline_query_current_chat"`
    Switch_inline_query_chosen_chat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`
    Copy_text *CopyTextButton `json:"copy_text"`
    Callback_game *CallbackGame `json:"callback_game"`
    Pay bool `json:"pay"`
}