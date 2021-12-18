package controllers

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"handle-big-post-request/queue"
	"handle-big-post-request/queue/payload"
	"handle-big-post-request/validate"
)

// Take note: our UserHandler has a UserStore injected in!
// Global data stores can be problematic long term.
type PostDataController struct {
	//PostData models.PostData
	//Sessions validate.UserPostFormUpload
	Ctx *fasthttp.RequestCtx
}

func (controller *PostDataController) Create(poolJob queue.PoolJob) {
	// get data request
	dataRequest, ok := controller.Ctx.MultipartForm()
	if ok != nil {
		fmt.Println("error param post")
		responseError(controller.Ctx, 400, "wrong format")
	}

	email, found := dataRequest.Value["email"]
	if !found {
		fmt.Println("error param post")
		responseError(controller.Ctx, 400, "missing param email")
	}

	name, found := dataRequest.Value["name"]
	if !found {
		fmt.Println("error param post")
		responseError(controller.Ctx, 400, "missing param name")
	}

	//validate
	valid := validate.UserPostFormUpload{email[0], name[0]}
	ok = valid.Validate()
	if ok != nil {
		fmt.Println("error param post")
		responseError(controller.Ctx, 400, "validate error")
	}

	// push to queue
	// create job data
	//jobData := queue.Job{payload.Payload{"test"}}

	// Push the work onto the queue.
	//log.Println("start push to queue")
	poolJob.PushDataToQueue(payload.Payload{"test"})
	//log.Println("end  push to queue")

	//response
	controller.Ctx.SetStatusCode(201)
	fmt.Fprintf(controller.Ctx, `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
}
