package queue

import (
	"fmt"
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
	fmt.Println("start Push job to Queue")
	p.Pool <- job
	fmt.Println("end Push job to Queue")
}

func (p *PoolJob) PushDataToQueue(payload payload.Payload) {
	jobData := Job{payload}
	p.PushJobToQueue(jobData)
}
