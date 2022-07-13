package concurrency

func New(maxGoRoutines int) *Context {
	c := Context{
		max:       maxGoRoutines,
		doneCh:    make(chan bool),
		allDoneCh: make(chan bool),
	}
	if maxGoRoutines != -1 {
		c.managerCh = make(chan any, maxGoRoutines)
		for i := 0; i < c.max; i++ {
			c.managerCh <- nil
		}
		go c.controller()
	}
	return &c
}
