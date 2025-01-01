package updates

type BusinessOpeningHours struct {
    Time_zone_name string `json:"time_zone_name"`
    Opening_hours *[]BusinessOpeningHoursInterval `json:"opening_hours"`
}