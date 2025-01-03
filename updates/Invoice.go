package updates

type Invoice struct {
    Title string `json:"title"`
    Description string `json:"description"`
    Start_parameter string `json:"start_parameter"`
    Currency string `json:"currency"`
    Total_amount int64 `json:"total_amount"`
}