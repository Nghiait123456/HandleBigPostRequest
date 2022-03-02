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
	addr = flag.String("addr", ":8080", "TCP address to listen to")
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
	controller := controllers.PostDataController{ctx}
	controller.Create(poolWorkerUpload)
}
