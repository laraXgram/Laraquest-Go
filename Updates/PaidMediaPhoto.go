package updates

type PaidMediaPhoto struct {
    Type string `json:"type"`
    Photo *[]PhotoSize `json:"photo"`
}