package updates

type ShippingOption struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Prices *[]LabeledPrice `json:"prices"`
}