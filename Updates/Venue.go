package updates

type Venue struct {
    Location *Location `json:"location"`
    Title string `json:"title"`
    Address string `json:"address"`
    Foursquare_id string `json:"foursquare_id"`
    Foursquare_type string `json:"foursquare_type"`
    Google_place_id string `json:"google_place_id"`
    Google_place_type string `json:"google_place_type"`
}