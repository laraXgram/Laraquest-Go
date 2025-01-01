package updates

type AffiliateInfo struct {
    Affiliate_user *User `json:"affiliate_user"`
    Affiliate_chat *Chat `json:"affiliate_chat"`
    Commission_per_mille int64 `json:"commission_per_mille"`
    Amount int64 `json:"amount"`
    Nanostar_amount int64 `json:"nanostar_amount"`
}