package updates

type UserProfilePhotos struct {
    Total_count int64 `json:"total_count"`
    Photos *[][]PhotoSize `json:"photos"`
}