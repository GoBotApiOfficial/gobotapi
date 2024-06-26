// Code AutoGenerated; DO NOT EDIT.

package types

// CallbackQuery Represents an incoming callback query from a callback button in an inline keyboard
// If the button that originated the query was attached to a message sent by the bot, the field message will be present
// If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present
// Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	ChatInstance    string                    `json:"chat_instance"`
	Data            string                    `json:"data,omitempty"`
	From            User                      `json:"from"`
	GameShortName   string                    `json:"game_short_name,omitempty"`
	ID              string                    `json:"id"`
	InlineMessageID string                    `json:"inline_message_id,omitempty"`
	Message         *MaybeInaccessibleMessage `json:"message,omitempty"`
}
