package updates

type LabeledPrice struct {
    Label string `json:"label"`
    Amount int64 `json:"amount"`
}