package updates

type MaskPosition struct {
    Point string `json:"point"`
    X_shift float64 `json:"x_shift"`
    Y_shift float64 `json:"y_shift"`
    Scale float64 `json:"scale"`
}