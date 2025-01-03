package updates

type PreCheckoutQuery struct {
    Id string `json:"id"`
    From *User `json:"from"`
    Currency string `json:"currency"`
    Total_amount int64 `json:"total_amount"`
    Invoice_payload string `json:"invoice_payload"`
    Shipping_option_id string `json:"shipping_option_id"`
    Order_info *OrderInfo `json:"order_info"`
}