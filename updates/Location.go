package updates

type Location struct {
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    Horizontal_accuracy float64 `json:"horizontal_accuracy"`
    Live_period int64 `json:"live_period"`
    Heading int64 `json:"heading"`
    Proximity_alert_radius int64 `json:"proximity_alert_radius"`
}