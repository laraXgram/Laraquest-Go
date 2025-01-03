package updates

type Gift struct {
    Id string `json:"id"`
    Sticker *Sticker `json:"sticker"`
    Star_count int64 `json:"star_count"`
    Upgrade_star_count int64 `json:"upgrade_star_count"`
    Total_count int64 `json:"total_count"`
    Remaining_count int64 `json:"remaining_count"`
}