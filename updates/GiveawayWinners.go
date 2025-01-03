package updates

type GiveawayWinners struct {
    Chat *Chat `json:"chat"`
    Giveaway_message_id int64 `json:"giveaway_message_id"`
    Winners_selection_date int64 `json:"winners_selection_date"`
    Winner_count int64 `json:"winner_count"`
    Winners *[]User `json:"winners"`
    Additional_chat_count int64 `json:"additional_chat_count"`
    Prize_star_count int64 `json:"prize_star_count"`
    Premium_subscription_month_count int64 `json:"premium_subscription_month_count"`
    Unclaimed_prize_count int64 `json:"unclaimed_prize_count"`
    Only_new_members bool `json:"only_new_members"`
    Was_refunded bool `json:"was_refunded"`
    Prize_description string `json:"prize_description"`
}