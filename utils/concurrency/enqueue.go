package concurrency

func (p *Pool) Enqueue(job func(params ...any), params ...any) {
	if p.jobs == 0 {
		go job(params...)
		return
	}
	p.queue <- struct{}{}
	go func() {
		defer func() {
			<-p.queue
		}()
		job(params...)
	}()
}
