package types

type ProximityAlertTriggered struct {
	Distance int `json:"distance"`
	Traveler User `json:"traveler"`
	Watcher User `json:"watcher"`
}
