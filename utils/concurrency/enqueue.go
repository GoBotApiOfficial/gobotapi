package concurrency

func (p *Pool) Enqueue(job func()) {
	if p.jobs == 0 {
		go job()
		return
	}
	p.queue <- struct{}{}
	go func() {
		defer func() {
			<-p.queue
		}()
		job()
	}()
}
