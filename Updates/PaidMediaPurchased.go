package updates

type PaidMediaPurchased struct {
    From *User `json:"from"`
    Paid_media_payload string `json:"paid_media_payload"`
}