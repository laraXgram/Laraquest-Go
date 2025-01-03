package updates

type ShippingAddress struct {
    Country_code string `json:"country_code"`
    State string `json:"state"`
    City string `json:"city"`
    Street_line1 string `json:"street_line1"`
    Street_line2 string `json:"street_line2"`
    Post_code string `json:"post_code"`
}