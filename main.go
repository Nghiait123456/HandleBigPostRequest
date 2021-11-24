package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

var (
	addr     = flag.String("addr", ":8080", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

var poolWorkerUpload = PoolJob{make(chan Job, 200000), 1500}

func main() {
	poolWorkerUpload.initQueue()
	fasthttp.ListenAndServe(":8090", requestHandler)
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Path()) {
	case "/user/nghia":
		uploadHanlde(ctx)
	default:
		ctx.Error("not found", fasthttp.StatusNotFound)
	}
}

func uploadHanlde(ctx *fasthttp.RequestCtx) {
	// create job data
	jobData := Job{Payload{"test"}}

	// Push the work onto the queue.
	log.Println("start push to queue")
	poolWorkerUpload.PushJobToQueue(jobData)
	log.Println("end  push to queue")

	//response
	ctx.SetStatusCode(201)
	fmt.Fprintf(ctx, `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
}
