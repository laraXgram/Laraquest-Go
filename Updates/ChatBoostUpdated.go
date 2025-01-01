package updates

type ChatBoostUpdated struct {
    Chat *Chat `json:"chat"`
    Boost *ChatBoost `json:"boost"`
}