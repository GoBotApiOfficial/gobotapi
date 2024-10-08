// Code AutoGenerated; DO NOT EDIT.

package types

// Update Represents an incoming update.At most one of the optional parameters can be present in any given update.
type Update struct {
	BusinessConnection      *BusinessConnection          `json:"business_connection,omitempty"`
	BusinessMessage         *Message                     `json:"business_message,omitempty"`
	CallbackQuery           *CallbackQuery               `json:"callback_query,omitempty"`
	ChannelPost             *Message                     `json:"channel_post,omitempty"`
	ChatBoost               *ChatBoostUpdated            `json:"chat_boost,omitempty"`
	ChatJoinRequest         *ChatJoinRequest             `json:"chat_join_request,omitempty"`
	ChatMember              *ChatMemberUpdated           `json:"chat_member,omitempty"`
	ChosenInlineResult      *ChosenInlineResult          `json:"chosen_inline_result,omitempty"`
	DeletedBusinessMessages *BusinessMessagesDeleted     `json:"deleted_business_messages,omitempty"`
	EditedBusinessMessage   *Message                     `json:"edited_business_message,omitempty"`
	EditedChannelPost       *Message                     `json:"edited_channel_post,omitempty"`
	EditedMessage           *Message                     `json:"edited_message,omitempty"`
	InlineQuery             *InlineQuery                 `json:"inline_query,omitempty"`
	Message                 *Message                     `json:"message,omitempty"`
	MessageReaction         *MessageReactionUpdated      `json:"message_reaction,omitempty"`
	MessageReactionCount    *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	MyChatMember            *ChatMemberUpdated           `json:"my_chat_member,omitempty"`
	Poll                    *Poll                        `json:"poll,omitempty"`
	PollAnswer              *PollAnswer                  `json:"poll_answer,omitempty"`
	PreCheckoutQuery        *PreCheckoutQuery            `json:"pre_checkout_query,omitempty"`
	PurchasedPaidMedia      *PaidMediaPurchased          `json:"purchased_paid_media,omitempty"`
	RemovedChatBoost        *ChatBoostRemoved            `json:"removed_chat_boost,omitempty"`
	ShippingQuery           *ShippingQuery               `json:"shipping_query,omitempty"`
	UpdateID                int64                        `json:"update_id"`
}
