package gobotapi

import (
	"fmt"
	"github.com/Squirrel-Network/gobotapi/types"
	"regexp"
	"strings"
)

func (ctx *BasicClient) OnCommand(command string, aliasList []string, handler func(client *Client, update types.Message)) {
	if len(aliasList) == 0 {
		aliasList = []string{"/"}
	}
	var prefixes []string
	for _, alias := range aliasList {
		prefixes = append(prefixes, fmt.Sprintf("(?:%s)", regexp.QuoteMeta(alias)))
	}
	var withoutPrefixCompiler *regexp.Regexp
	cmdHandler := func(client *Client, message types.Message) {
		if withoutPrefixCompiler == nil {
			withoutPrefixCompiler, _ = regexp.Compile(fmt.Sprintf("(?i)((%s)%s(?:@?%s)?)(?:\\s|$)", strings.Join(prefixes, "|"), command, client.me.Username))
		}
		text := message.Text
		if len(text) == 0 {
			text = message.Caption
		}
		if len(text) > 0 {
			matches := withoutPrefixCompiler.FindAllStringSubmatch(text, -1)
			if len(matches) == 0 || !strings.HasPrefix(text, matches[0][1]) {
				return
			}
			handler(client, message)
		}
	}
	ctx.OnMessage(cmdHandler)
	ctx.OnEditedMessage(cmdHandler)
}

func (ctx *BasicClient) OnCallbackQueryData(data string, handler func(client *Client, update types.CallbackQuery)) {
	ctx.OnCallbackQuery(func(client *Client, update types.CallbackQuery) {
		if update.Data == data {
			handler(client, update)
		}
	})
}
