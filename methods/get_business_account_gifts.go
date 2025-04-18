// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/GoBotApiOfficial/gobotapi/types"
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
)

// GetBusinessAccountGifts Returns the gifts received and owned by a managed business account
// Requires the can_view_gifts_and_stars business bot right
// Returns OwnedGifts on success.
type GetBusinessAccountGifts struct {
	BusinessConnectionID string `json:"business_connection_id"`
	ExcludeLimited       bool   `json:"exclude_limited,omitempty"`
	ExcludeSaved         bool   `json:"exclude_saved,omitempty"`
	ExcludeUnique        bool   `json:"exclude_unique,omitempty"`
	ExcludeUnlimited     bool   `json:"exclude_unlimited,omitempty"`
	ExcludeUnsaved       bool   `json:"exclude_unsaved,omitempty"`
	Limit                int    `json:"limit,omitempty"`
	Offset               string `json:"offset,omitempty"`
	SortByPrice          bool   `json:"sort_by_price,omitempty"`
}

func (entity *GetBusinessAccountGifts) ProgressCallable() rawTypes.ProgressCallable {
	return nil
}

func (entity *GetBusinessAccountGifts) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetBusinessAccountGifts) MethodName() string {
	return "getBusinessAccountGifts"
}

func (GetBusinessAccountGifts) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.OwnedGifts `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result{
		Kind:   types.TypeOwnedGifts,
		Result: x1.Result,
	}
	return &result, nil
}
