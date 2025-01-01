package updates

type TransactionPartnerUser struct {
    Type string `json:"type"`
    User *User `json:"user"`
    Affiliate *AffiliateInfo `json:"affiliate"`
    Invoice_payload string `json:"invoice_payload"`
    Subscription_period int64 `json:"subscription_period"`
    Paid_media *[]PaidMedia `json:"paid_media"`
    Paid_media_payload string `json:"paid_media_payload"`
    Gift *Gift `json:"gift"`
}