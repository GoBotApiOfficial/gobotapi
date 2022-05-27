package types

type ReplyKeyboardMarkup struct {
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Keyboard [][]KeyboardButton `json:"keyboard,omitempty"`
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	Selective bool `json:"selective,omitempty"`
}
