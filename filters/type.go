// Code AutoGenerated; DO NOT EDIT.

package filters

import (
	"github.com/GoBotApiOfficial/gobotapi"
	"github.com/GoBotApiOfficial/gobotapi/types"
)

type DataFilter struct {
	Chat      *types.Chat
	Date      *int64
	From      *types.User
	Message   *types.MaybeInaccessibleMessage
	Client    *gobotapi.Client
	RawUpdate any
}
