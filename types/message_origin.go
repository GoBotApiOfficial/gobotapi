// Code AutoGenerated; DO NOT EDIT.

package types

// MessageOrigin This object describes the origin of a message
// It can be one of
//   - MessageOriginUser
//   - MessageOriginHiddenUser
//   - MessageOriginChat
//   - MessageOriginChannel
type MessageOrigin struct {
	AuthorSignature string `json:"author_signature"`
	Chat            Chat   `json:"chat"`
	Date            int64  `json:"date"`
	MessageID       int64  `json:"message_id"`
	SenderChat      Chat   `json:"sender_chat"`
	SenderUser      User   `json:"sender_user"`
	SenderUserName  string `json:"sender_user_name"`
	Type            string `json:"type"`
}

func (x MessageOrigin) Kind() int {
	switch x.Type {
	case "user":
		return TypeMessageOriginUser
	case "hidden_user":
		return TypeMessageOriginHiddenUser
	case "chat":
		return TypeMessageOriginChat
	case "channel":
		return TypeMessageOriginChannel
	default:
		return -1
	}
}
