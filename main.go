package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"handle-big-post-request/controllers"
	"handle-big-post-request/queue"
	"log"
)

var poolWorkerUpload = queue.PoolJob{make(chan queue.Job, 100000), 15000}

var (
	addr = flag.String("addr", ":8091", "TCP address to listen to")
)

func main() {
	poolWorkerUpload.InitQueue()
	fmt.Println("init queue succcess")
	if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/postFrom":
		uploadHanlde(ctx)
	default:
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
}

func uploadHanlde(ctx *fasthttp.RequestCtx) {
	//dataRequest, _ := ctx.MultipartForm()
	//fmt.Println("post", dataRequest, dataRequest.Value["test"][0])
	controller := controllers.PostDataController{ctx}
	controller.Create(poolWorkerUpload)

	//// create job data
	//jobData := Job{Payload{"test"}}
	//
	//// Push the work onto the queue.
	////log.Println("start push to queue")
	//poolWorkerUpload.PushJobToQueue(jobData)
	//log.Println("end  push to queue")

	//response
	//ctx.SetStatusCode(201)
	//fmt.Fprintf(ctx, `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
}

//func pushDataToQueue(payload Payload) {
//	jobData := Job{payload}
//	poolWorkerUpload.PushJobToQueue(jobData)
//}
