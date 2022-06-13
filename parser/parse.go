package parser

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
)

func Parse(data string, entities []types.MessageEntity) string {
	text := []rune(data)
	openingTags := make([]string, len(text))
	closingTags := make([]string, len(text))
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
	var newText string
	for i, t := range text {
		newText += openingTags[i] + string(t) + closingTags[i]
	}
	return newText
}
