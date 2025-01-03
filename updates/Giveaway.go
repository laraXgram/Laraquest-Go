package updates

type Giveaway struct {
    Chats *[]Chat `json:"chats"`
    Winners_selection_date int64 `json:"winners_selection_date"`
    Winner_count int64 `json:"winner_count"`
    Only_new_members bool `json:"only_new_members"`
    Has_public_winners bool `json:"has_public_winners"`
    Prize_description string `json:"prize_description"`
    Country_codes []string `json:"country_codes"`
    Prize_star_count int64 `json:"prize_star_count"`
    Premium_subscription_month_count int64 `json:"premium_subscription_month_count"`
}