package concurrency

func New(maxGoRoutines int) *Context {
	c := Context{
		max:       maxGoRoutines,
		managerCh: make(chan any, maxGoRoutines),
		doneCh:    make(chan bool),
		allDoneCh: make(chan bool),
	}
	for i := 0; i < c.max; i++ {
		c.managerCh <- nil
	}
	go c.controller()
	return &c
}
