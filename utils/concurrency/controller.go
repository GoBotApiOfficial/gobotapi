package concurrency

func (c *Context) controller() {
	for {
		<-c.doneCh
		c.managerCh <- nil
		if c.closed == true && c.runningCount == 0 {
			break
		}
	}
	c.allDoneCh <- true
}
