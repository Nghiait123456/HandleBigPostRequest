package queue

import (
	"handle-big-post-request/queue/payload"
)

// Job represents the job to be run
type Job struct {
	PayloadJob payload.Payload
}

// all data job
type PoolJob struct {
	Pool      chan Job
	MaxWorker int
}

func (p *PoolJob) InitQueue() {
	for i := 0; i < p.MaxWorker; i++ {
		p.StartOneWorker()
	}
}

func (p *PoolJob) StartOneWorker() {
	go func() {
		for {
			// get job from pool job
			job := <-p.Pool
			job.PayloadJob.Handle()
		}
	}()
}

func (p *PoolJob) PushJobToQueue(job Job) {
	p.Pool <- job
}

func (p *PoolJob) PushDataToQueue(payload payload.Payload) {
	jobData := Job{payload}
	p.PushJobToQueue(jobData)
}
