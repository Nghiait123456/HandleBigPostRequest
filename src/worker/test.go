package main

import (
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/kataras/iris/v12"
	"handle-big-post-request/src/worker/w_pool_local_in_mm"
)

var wp *workerpool.WorkerPool

func init() {
	w_pool_local_in_mm.ResignAllWorker()
}

func main() {
	//fmt.Println("start new worker pool")
	//wp = workerpool.New(2)
	//fmt.Println("start push job to pool")
	//requests := []string{"alpha", "beta", "gamma", "delta", "epsilon"}
	//todo example call multil api and response
	// tham khao Iris call api
	// push job anysnc
	go func() {
		for i := 0; i < 1000; i++ {
			r := "test"
			w_pool_local_in_mm.WorkerNormal().PushJobToPollWork(func() {
				fmt.Println("Handling request:", r)
			})
		}
	}()

	app := iris.New()
	app.Get("/addJob", func(ctx iris.Context) {
		r := "addJob"
		wp.Submit(func() {
			fmt.Println("Handling request:", r)
		})
		ctx.WriteString("pong")
	})

	if err := app.Listen(":8080", iris.WithoutBanner); err != nil {
		app.Logger().Warn("Shutdown with error: " + err.Error())
	}

	//fmt.Println("start stopWait")
	//wp.StopWait()
	//fmt.Println("end Stop wait")
}
