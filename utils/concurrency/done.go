package concurrency

import "sync/atomic"

func (c *Context) Done() {
	if c.max == -1 {
		return
	}
	atomic.AddInt64(&c.runningCount, -1)
	c.doneCh <- true
}
