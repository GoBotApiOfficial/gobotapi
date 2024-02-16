// Code AutoGenerated; DO NOT EDIT.

package gobotapi

import (
	"errors"
	"gobotapi/types"
	rawTypes "gobotapi/types/raw"
)

func (ctx *Client) DownloadMedia(message types.Message, filePath string, progress rawTypes.ProgressCallable) error {
	if message.Animation != nil {
		return ctx.DownloadFile(message.Animation.FileID, filePath, progress)
	}
	if message.Audio != nil {
		return ctx.DownloadFile(message.Audio.FileID, filePath, progress)
	}
	if message.Document != nil {
		return ctx.DownloadFile(message.Document.FileID, filePath, progress)
	}
	if len(message.Photo) > 0 {
		var bestQuality types.PhotoSize
		for _, file := range message.Photo {
			if file.Width > bestQuality.Width {
				bestQuality = file
			}
		}
		return ctx.DownloadFile(bestQuality.FileID, filePath, progress)
	}
	if message.Sticker != nil {
		return ctx.DownloadFile(message.Sticker.FileID, filePath, progress)
	}
	if message.Video != nil {
		return ctx.DownloadFile(message.Video.FileID, filePath, progress)
	}
	if message.VideoNote != nil {
		return ctx.DownloadFile(message.VideoNote.FileID, filePath, progress)
	}
	if message.Voice != nil {
		return ctx.DownloadFile(message.Voice.FileID, filePath, progress)
	}
	if len(message.NewChatPhoto) > 0 {
		var bestQuality types.PhotoSize
		for _, file := range message.NewChatPhoto {
			if file.Width > bestQuality.Width {
				bestQuality = file
			}
		}
		return ctx.DownloadFile(bestQuality.FileID, filePath, progress)
	}
	return errors.New("no files found")
}
