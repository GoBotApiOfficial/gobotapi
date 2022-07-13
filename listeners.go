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
	withoutPrefixCompilers := make(map[int64]*regexp.Regexp)
	cmdHandler := func(client *Client, message types.Message) {
		if withoutPrefixCompilers[client.me.ID] == nil {
			withoutPrefixCompilers[client.me.ID], _ = regexp.Compile(fmt.Sprintf("(?i)((%s)%s(?:@?%s)?)(?:\\s|$)", strings.Join(prefixes, "|"), command, client.me.Username))
		}
		text := message.Text
		if len(text) == 0 {
			text = message.Caption
		}
		if len(text) > 0 {
			matches := withoutPrefixCompilers[client.me.ID].FindAllStringSubmatch(text, -1)
			if len(matches) == 0 || !strings.HasPrefix(text, matches[0][1]) {
				return
			}
			handler(client, message)
		}
	}
	ctx.OnMessage(cmdHandler)
	ctx.OnEditedMessage(cmdHandler)
}
