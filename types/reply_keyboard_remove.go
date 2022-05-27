package types

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective bool `json:"selective,omitempty"`
}
