package updates

type ChatInviteLink struct {
    Invite_link string `json:"invite_link"`
    Creator *User `json:"creator"`
    Creates_join_request bool `json:"creates_join_request"`
    Is_primary bool `json:"is_primary"`
    Is_revoked bool `json:"is_revoked"`
    Name string `json:"name"`
    Expire_date int64 `json:"expire_date"`
    Member_limit int64 `json:"member_limit"`
    Pending_join_request_count int64 `json:"pending_join_request_count"`
    Subscription_period int64 `json:"subscription_period"`
    Subscription_price int64 `json:"subscription_price"`
}