package updates

type SuccessfulPayment struct {
    Currency string `json:"currency"`
    Total_amount int64 `json:"total_amount"`
    Invoice_payload string `json:"invoice_payload"`
    Subscription_expiration_date int64 `json:"subscription_expiration_date"`
    Is_recurring bool `json:"is_recurring"`
    Is_first_recurring bool `json:"is_first_recurring"`
    Shipping_option_id string `json:"shipping_option_id"`
    Order_info *OrderInfo `json:"order_info"`
    Telegram_payment_charge_id string `json:"telegram_payment_charge_id"`
    Provider_payment_charge_id string `json:"provider_payment_charge_id"`
}