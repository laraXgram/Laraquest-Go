package updates

type RefundedPayment struct {
    Currency string `json:"currency"`
    Total_amount int64 `json:"total_amount"`
    Invoice_payload string `json:"invoice_payload"`
    Telegram_payment_charge_id string `json:"telegram_payment_charge_id"`
    Provider_payment_charge_id string `json:"provider_payment_charge_id"`
}