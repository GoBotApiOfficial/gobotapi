// Code AutoGenerated; DO NOT EDIT.

package gobotapi

import "github.com/Squirrel-Network/gobotapi/types"


func (ctx *Client) DownloadMessage(message types.Message, filePath string) error {
	if message.Animation != nil {
		return ctx.DownloadFile(message.Animation.FileID, filePath)
	}
	if message.Audio != nil {
		return ctx.DownloadFile(message.Audio.FileID, filePath)
	}
	if message.Document != nil {
		return ctx.DownloadFile(message.Document.FileID, filePath)
	}
	if message.Sticker != nil {
		return ctx.DownloadFile(message.Sticker.FileID, filePath)
	}
	if message.Video != nil {
		return ctx.DownloadFile(message.Video.FileID, filePath)
	}
	if message.VideoNote != nil {
		return ctx.DownloadFile(message.VideoNote.FileID, filePath)
	}
	if message.Voice != nil {
		return ctx.DownloadFile(message.Voice.FileID, filePath)
	}
	return nil
}