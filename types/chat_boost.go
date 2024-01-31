// Code AutoGenerated; DO NOT EDIT.

package types

// ChatBoost This object contains information about a chat boost.
type ChatBoost struct {
	AddDate        int64           `json:"add_date"`
	BoostID        string          `json:"boost_id"`
	ExpirationDate int64           `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}
