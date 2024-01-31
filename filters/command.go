package filters

import (
	"fmt"
	"github.com/GoBotApiOfficial/gobotapi/types"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

func Command(command string, aliasList ...string) FilterOperand {
	if len(aliasList) == 0 {
		aliasList = []string{"/"}
	}
	var prefixes []string
	for _, alias := range aliasList {
		prefixes = append(prefixes, fmt.Sprintf("(?:%s)", regexp.QuoteMeta(alias)))
	}
	withoutPrefixCompilers := make(map[int64]*regexp.Regexp)
	var mutex sync.RWMutex
	return func(values *DataFilter) bool {
		if reflect.TypeOf(values.RawUpdate).String() == "types.Message" {
			message := values.RawUpdate.(types.Message)
			mutex.Lock()
			botInfo := values.Client.Self()
			if withoutPrefixCompilers[botInfo.ID] == nil {
				withoutPrefixCompilers[botInfo.ID], _ = regexp.Compile(fmt.Sprintf("(?i)((%s)%s(?:@?%s)?)(?:\\s|$)", strings.Join(prefixes, "|"), command, botInfo.Username))
			}
			mutex.Unlock()
			text := message.Text
			if len(text) == 0 {
				text = message.Caption
			}
			if len(text) > 0 {
				mutex.Lock()
				matches := withoutPrefixCompilers[botInfo.ID].FindAllStringSubmatch(text, -1)
				mutex.Unlock()
				return len(matches) > 0 && strings.HasPrefix(text, matches[0][1])
			}
		}
		return false
	}
}
