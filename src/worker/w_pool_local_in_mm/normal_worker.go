package w_pool_local_in_mm

import (
	"github.com/gammazero/workerpool"
)

var nWorker *normalWorker

// normalWorker is pool worker of normal job, all job in action every day: call api, send ipn, calculator,...
// it's simple and very simple use
// func newNormalWorkerPool(maxWorker int) => start and init workPool
// func w_pool_local_in_mm.WorkerNormal().PushJobToPollWork(func() {
//				fmt.Println("Handling request:", r)
//			})
// ==> push jon to wp, use variable in local function, it's very
//    Convenience, using variables in the current code flow,
//    in this current fc, no need to initialize additional objects,   => seamless feeling when code
type normalWorker struct {
	wp        *workerpool.WorkerPool
	maxWorker int
}

func (wk *normalWorker) PushJobToPollWork(task func()) {
	wk.wp.Submit(task)
}

func WorkerNormal() *normalWorker {
	if nWorker == nil {
		panic("please init normal worker before use")
	}

	return nWorker
}

func newNormalWorkerPool(maxWorker int) *normalWorker {
	wp := workerpool.New(maxWorker)
	nWorker = &normalWorker{
		wp,
		maxWorker,
	}
	return nWorker
}
