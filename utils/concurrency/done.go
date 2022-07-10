package concurrency

import "sync/atomic"

func (c *Context) Done() {
	atomic.AddInt64(&c.runningCount, -1)
	c.doneCh <- true
}
