package parser

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	"unicode/utf16"
)

func Parse(data string, entities []types.MessageEntity) string {
	text := utf16.Encode([]rune(data))
	openingTags := make([]string, len(text)+1)
	closingTags := make([]string, len(text)+1)
	for _, entity := range entities {
		name := HtmlTag(entity.Type)
		if len(name) == 0 {
			continue
		}
		start := entity.Offset
		end := start + entity.Length - 1
		var startTag, endTag string
		switch name {
		case "a":
			if entity.User != nil {
				startTag = fmt.Sprintf("<%s href='tg://user?id=%d'>", name, entity.User.ID)
			} else {
				startTag = fmt.Sprintf("<%s href='%s'>", name, entity.URL)
			}
			endTag = fmt.Sprintf("</%s>", name)
		default:
			startTag = fmt.Sprintf("<%s>", name)
			endTag = fmt.Sprintf("</%s>", name)
			if len(entity.Language) > 0 {
				startTag += fmt.Sprintf("<code class=\"language-%s\">", entity.Language)
				endTag = "</code>" + endTag
			}
		}
		openingTags[start] += startTag
		closingTags[end] = endTag + closingTags[end]
	}
	var newText []uint16
	for i, t := range text {
		newText = append(newText, utf16.Encode([]rune(openingTags[i]))...)
		if t == 60 {
			newText = append(newText, utf16.Encode([]rune("&lt;"))...)
		} else if t == 62 {
			newText = append(newText, utf16.Encode([]rune("&gt;"))...)
		} else {
			newText = append(newText, t)
		}
		newText = append(newText, utf16.Encode([]rune(closingTags[i]))...)
	}
	return string(utf16.Decode(newText))
}
