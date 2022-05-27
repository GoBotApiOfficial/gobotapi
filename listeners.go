package gobotapi

import (
	"github.com/Squirrel-Network/gobotapi/types"
	"fmt"
	"regexp"
	"strings"
)

func (ctx *Client) OnCommand(command string, aliasList []string, handler func(update types.Message)) {
	if len(aliasList) == 0 {
		aliasList = []string{"/"}
	}
	r, _ := regexp.Compile(fmt.Sprintf("(?i)(%s(?:@?%s)?)(?:\\s|$)", command, ctx.botUsername))
	cmdHandler := func(message types.Message) {
		text := message.Text
		if len(text) == 0 {
			text = message.Caption
		}
		if len(text) > 0 {
			for _, alias := range aliasList {
				if !strings.HasPrefix(message.Text, alias) {
					continue
				}
				withoutPrefix := strings.TrimPrefix(text, alias)
				matches := r.FindAllStringSubmatch(withoutPrefix, -1)
				if len(matches) == 0 {
					continue
				}
				handler(message)
				break
			}
		}
	}
	ctx.OnMessage(cmdHandler)
	ctx.OnEditedMessage(cmdHandler)
}

