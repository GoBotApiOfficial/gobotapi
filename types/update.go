package types

// Update Represents an incoming update.At most one of the optional parameters can be present in any given update.
type Update struct {
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	ChannelPost *Message `json:"channel_post,omitempty"`
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	EditedMessage *Message `json:"edited_message,omitempty"`
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	Message *Message `json:"message,omitempty"`
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	Poll *Poll `json:"poll,omitempty"`
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	UpdateId int64 `json:"update_id"`
}
