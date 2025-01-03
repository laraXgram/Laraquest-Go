package updates

type KeyboardButtonRequestUsers struct {
    Request_id int64 `json:"request_id"`
    User_is_bot bool `json:"user_is_bot"`
    User_is_premium bool `json:"user_is_premium"`
    Max_quantity int64 `json:"max_quantity"`
    Request_name bool `json:"request_name"`
    Request_username bool `json:"request_username"`
    Request_photo bool `json:"request_photo"`
}