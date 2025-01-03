package updates

type GiveawayCompleted struct {
    Winner_count int64 `json:"winner_count"`
    Unclaimed_prize_count int64 `json:"unclaimed_prize_count"`
    Giveaway_message *Message `json:"giveaway_message"`
    Is_star_giveaway bool `json:"is_star_giveaway"`
}