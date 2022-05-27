package raw

import (
	"context"
	"time"
)

type CancelableContext struct {
	Id     int64
	Cancel context.CancelFunc
}

func (ctx *CancelableContext) GenerateId() {
	ctx.Id = time.Now().UnixNano()
}
