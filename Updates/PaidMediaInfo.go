package updates

type PaidMediaInfo struct {
    Star_count int64 `json:"star_count"`
    Paid_media *[]PaidMedia `json:"paid_media"`
}