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

	// Disable Console Color, you don't need console color when writing the logs to file.
	//gin.DisableConsoleColor()
	//r := gin.Default()
	////
	////// Logging to a file.
	////f, _ := os.Create("gin.log")
	////gin.DefaultWriter = io.MultiWriter(f)
	////
	////// init pool

	poolWorkerUpload.InitQueue()

	//r.GET("/upload", func(c *gin.Context) {
	//	jobData := Job{Payload{"test"}}
	//
	//	// Push the work onto the queue.
	//	log.Println("start push to queue")
	//	pool.Pool <- jobData
	//	log.Println("end  push to queue")
	//
	//	c.JSON(200, gin.H{
	//		"message": "upload success",
	//	})
	//})
	//
	//fmt.Print("start run serve ")
	//r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//fmt.Print("end run serve")

	//m := &http.ServeMux{}
	//m.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
	//	fmt.Fprintf(writer, "Hello, %q", html.EscapeString(request.URL.Path))
	//})
	//
	//http.ListenAndServe(":8080", m)

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
	// push job to queue
	jobData := Job{Payload{"test"}}

	// Push the work onto the queue.
	log.Println("start push to queue")
	poolWorkerUpload.Pool <- jobData
	log.Println("end  push to queue")

	ctx.SetStatusCode(201)
	fmt.Fprintf(ctx, `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
}
