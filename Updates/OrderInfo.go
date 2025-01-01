package updates

type OrderInfo struct {
    Name string `json:"name"`
    Phone_number string `json:"phone_number"`
    Email string `json:"email"`
    Shipping_address *ShippingAddress `json:"shipping_address"`
}