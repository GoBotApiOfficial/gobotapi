package types

type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective bool `json:"selective,omitempty"`
}
