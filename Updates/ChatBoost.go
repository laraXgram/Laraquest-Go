package updates

type ChatBoost struct {
    Boost_id string `json:"boost_id"`
    Add_date int64 `json:"add_date"`
    Expiration_date int64 `json:"expiration_date"`
    Source *ChatBoostSource `json:"source"`
}