package updates

type ProximityAlertTriggered struct {
    Traveler *User `json:"traveler"`
    Watcher *User `json:"watcher"`
    Distance int64 `json:"distance"`
}