package updates

type GameHighScore struct {
    Position int64 `json:"position"`
    User *User `json:"user"`
    Score int64 `json:"score"`
}