package concurrency

func NewPool(size int) *Pool {
	if size < 0 {
		size = 0
	}
	return &Pool{
		jobs:  size,
		queue: make(chan struct{}, size),
	}
}
