package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"handle-big-post-request/config"
	"handle-big-post-request/controllers"
	"handle-big-post-request/handle"
	"handle-big-post-request/queue"
)

var poolWorkerUpload queue.PoolJob

func init() {
	poolWorkerUpload = queue.PoolJob{make(chan queue.Job, 200000), 80000}
	config.Init("./config.yml")
	// todo cmt for test on EC2
	//database.Init()
}

func main() {
	poolWorkerUpload.InitQueue()
	fmt.Println("init queue succcess")
	app := iris.New()

	postDataController := controllers.PostDataController{&poolWorkerUpload}
	postDataHandle := handle.PostData{
		app,
		&poolWorkerUpload,
		&postDataController,
	}

	handle.ResignRoutePostData(&postDataHandle)

	app.Listen(":8080")
}
