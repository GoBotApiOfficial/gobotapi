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
	var withoutPrefixCompiler *regexp.Regexp
	cmdHandler := func(client *Client, message types.Message) {
		if withoutPrefixCompiler == nil {
			withoutPrefixCompiler, _ = regexp.Compile(fmt.Sprintf("(?i)(%s(?:@?%s)?)(?:\\s|$)", command, client.me.Username))
		}
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
				matches := withoutPrefixCompiler.FindAllStringSubmatch(withoutPrefix, -1)
				if len(matches) == 0 {
					continue
				}
				if strings.HasPrefix(withoutPrefix, command) {
					handler(client, message)
					break
				}
			}
		}
	}
	ctx.OnMessage(cmdHandler)
	ctx.OnEditedMessage(cmdHandler)
}
