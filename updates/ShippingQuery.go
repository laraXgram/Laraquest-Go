package updates

type ShippingQuery struct {
    Id string `json:"id"`
    From *User `json:"from"`
    Invoice_payload string `json:"invoice_payload"`
    Shipping_address *ShippingAddress `json:"shipping_address"`
}