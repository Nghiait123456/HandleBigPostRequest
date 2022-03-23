package main

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"handle-big-post-request/config"
	"handle-big-post-request/controllers"
	"handle-big-post-request/handle"
	"handle-big-post-request/logsCustom"
	"handle-big-post-request/queue"
	"io"
	"os"
)

var poolWorkerUpload queue.PoolJob

func init() {
	poolWorkerUpload = queue.PoolJob{make(chan queue.Job, 400000), 150000}
	config.Init("./config.yml")
	// todo cmt for test on EC2
	//database.Init()
}

func main() {
	/**
	Init log file
	*/
	f := logsCustom.NewLogFile()
	defer f.Close()

	/**
	Init worker pool
	*/
	poolWorkerUpload.InitQueue()
	fmt.Println("init queue succcess")

	/**
	Init, config web server
	*/
	app := iris.New()
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	app.Logger().SetOutput(f)
	app.Logger().SetLevel("debug")

	/**
	Init config
	*/
	config.Init("./config.yml")

	/**
	global middleware
	*/

	/**
	global logs
	*/
	GlobalLogsCf := config.AllConfig.Logs
	GLogs := logsCustom.Global{GlobalLogsCf.IsLogGlobal, GlobalLogsCf.IsPrintGlobal, GlobalLogsCf.IsUseLogGlobalMode}
	GLogs.ResignGlobalLog(app)

	/**
	Resign List router
	*/

	/**
	 *post form data handle
	 */
	postDataController := controllers.PostDataController{&poolWorkerUpload}
	postDataHandle := handle.PostData{
		app,
		&poolWorkerUpload,
		&postDataController,
	}

	handle.ResignRoutePostData(&postDataHandle)

	//health check
	app.Get("/ping", func(ctx iris.Context) {
		ctx.WriteString("pong")
	})

	//start server
	if err := app.Listen(":8080", iris.WithoutBanner); err != nil {
		app.Logger().Warn("Shutdown with error: " + err.Error())
	}
}
