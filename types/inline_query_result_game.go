package types

import "encoding/json"

type InlineQueryResultGame struct {
	GameShortName string `json:"game_short_name"`
	Id string `json:"id"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (entity InlineQueryResultGame) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		GameShortName string `json:"game_short_name"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	} {
		Type: "game",
		Id: entity.Id,
		GameShortName: entity.GameShortName,
		ReplyMarkup: entity.ReplyMarkup,
	}
	return json.Marshal(alias)
}
