// Code AutoGenerated; DO NOT EDIT.

package gobotapi

import "github.com/Squirrel-Network/gobotapi/types"


func ContainsFiles(message types.Message) bool {
	return message.Animation != nil || 
		message.Audio != nil || 
		message.Document != nil || 
		message.Sticker != nil || 
		message.Video != nil || 
		message.VideoNote != nil || 
		message.Voice != nil
}
