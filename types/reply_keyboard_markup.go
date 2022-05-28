package types

// ReplyKeyboardMarkup Represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Keyboard [][]KeyboardButton `json:"keyboard,omitempty"`
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	Selective bool `json:"selective,omitempty"`
}
