package updates

type BusinessLocation struct {
    Address string `json:"address"`
    Location *Location `json:"location"`
}