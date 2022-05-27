package types

type Venue struct {
	Address string `json:"address"`
	FoursquareId string `json:"foursquare_id,omitempty"`
	FoursquareType string `json:"foursquare_type,omitempty"`
	GooglePlaceId string `json:"google_place_id,omitempty"`
	GooglePlaceType string `json:"google_place_type,omitempty"`
	Location Location `json:"location"`
	Title string `json:"title"`
}
