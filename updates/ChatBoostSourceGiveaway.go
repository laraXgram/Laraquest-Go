package updates

type ChatBoostSourceGiveaway struct {
    Source string `json:"source"`
    Giveaway_message_id int64 `json:"giveaway_message_id"`
    User *User `json:"user"`
    Prize_star_count int64 `json:"prize_star_count"`
    Is_unclaimed bool `json:"is_unclaimed"`
}