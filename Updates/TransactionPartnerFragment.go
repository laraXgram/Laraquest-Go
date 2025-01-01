package updates

type TransactionPartnerFragment struct {
    Type string `json:"type"`
    Withdrawal_state *RevenueWithdrawalState `json:"withdrawal_state"`
}