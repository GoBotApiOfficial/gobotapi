package concurrency

type Pool struct {
	jobs  int
	queue chan struct{}
}
