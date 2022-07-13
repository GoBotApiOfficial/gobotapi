package concurrency

import "sync/atomic"

func (c *Context) Wait() {
	if c.max == -1 {
		return
	}
	<-c.managerCh
	atomic.AddInt64(&c.runningCount, 1)
}
