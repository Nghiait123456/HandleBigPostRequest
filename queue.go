package main

// Job represents the job to be run
type Job struct {
	PayloadJob Payload
}

// all data job
type PoolJob struct {
	Pool      chan Job
	MaxWorker int
}

func (p *PoolJob) initQueue() {
	for i := 0; i < p.MaxWorker; i++ {
		p.StartOneWorker()
	}
}

func (p *PoolJob) StartOneWorker() {
	go func() {
		for {
			// get job from pool job
			//log.Println("get data from pool")
			//wait until have job in channel
			job := <-p.Pool
			//log.Println("have data, start handle job")
			job.PayloadJob.Handle()
			//log.Println("have data, end handle job")
		}
	}()
}

func (p *PoolJob) PushJobToQueue(job Job) {
	p.Pool <- job
}
