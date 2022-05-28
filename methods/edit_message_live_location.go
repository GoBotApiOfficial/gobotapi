package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

// EditMessageLiveLocation Use this method to edit live location messages
// A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocation struct {
	ChatId int64 `json:"chat_id,omitempty"`
	Heading int `json:"heading,omitempty"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	MessageId int64 `json:"message_id,omitempty"`
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (entity *EditMessageLiveLocation) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (EditMessageLiveLocation) MethodName() string {
	return "editMessageLiveLocation"
}

func (EditMessageLiveLocation) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x0 struct {
		Result bool `json:"result"`
	}
	_ = json.Unmarshal(response, &x0)
	if x0.Result {
		result := rawTypes.Result {
			Kind: types.TypeBoolean,
			Result: true,
		}
		return &result, nil
	} else {
		var x1 struct {
			Result types.Message `json:"result"`
		}
		err := json.Unmarshal(response, &x1)
		if err != nil {
			return nil, err
		}
		result := rawTypes.Result {
			Kind: types.TypeMessage,
			Result: x1.Result,
		}
		return &result, nil
	}
}
