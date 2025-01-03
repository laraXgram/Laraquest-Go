package updates

type TransactionPartnerAffiliateProgram struct {
    Type string `json:"type"`
    Sponsor_user *User `json:"sponsor_user"`
    Commission_per_mille int64 `json:"commission_per_mille"`
}