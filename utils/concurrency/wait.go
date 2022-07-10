package concurrency

import "sync/atomic"

func (c *Context) Wait() {
	<-c.managerCh
	atomic.AddInt64(&c.runningCount, 1)
}
