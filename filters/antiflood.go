package filters

import (
	rawTypes "github.com/GoBotApiOfficial/gobotapi/types/raw"
	"time"
)

func AntiFlood(numEvents int, inTime time.Duration, banTime time.Duration) FilterOperand {
	return func(values *DataFilter) bool {
		if values.Date != nil {
			var id int64
			if values.Message != nil {
				id = values.Message.Chat.ID
			} else if values.Chat != nil {
				id = values.Chat.ID
			} else if values.From != nil {
				id = values.From.ID
			}
			if values.Client.AntiFloodData[id] != nil {
				lastEvent := time.Unix(values.Client.AntiFloodData[id].LastEvent, 0)
				messageDate := time.Unix(*values.Date, 0)
				if values.Client.AntiFloodData[id].NumEvents >= numEvents && lastEvent.Add(banTime).Sub(messageDate) > 0 {
					return false
				} else if lastEvent.Add(inTime).Sub(messageDate) > 0 {
					values.Client.AntiFloodData[id].NumEvents++
					return true
				}
			}
			values.Client.AntiFloodData[values.Chat.ID] = &rawTypes.AntiFloodData{
				NumEvents: 1,
				LastEvent: *values.Date,
			}
		}
		return true
	}
}
