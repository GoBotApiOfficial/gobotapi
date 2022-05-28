package types

// InputVenueMessageContent Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Address string `json:"address"`
	FoursquareId string `json:"foursquare_id,omitempty"`
	FoursquareType string `json:"foursquare_type,omitempty"`
	GooglePlaceId string `json:"google_place_id,omitempty"`
	GooglePlaceType string `json:"google_place_type,omitempty"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Title string `json:"title"`
}
