package concurrency

type Context struct {
	max          int
	managerCh    chan interface{}
	doneCh       chan bool
	allDoneCh    chan bool
	closed       bool
	runningCount int64
}
