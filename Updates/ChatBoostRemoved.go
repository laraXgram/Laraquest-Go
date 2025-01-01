package updates

type ChatBoostRemoved struct {
    Chat *Chat `json:"chat"`
    Boost_id string `json:"boost_id"`
    Remove_date int64 `json:"remove_date"`
    Source *ChatBoostSource `json:"source"`
}