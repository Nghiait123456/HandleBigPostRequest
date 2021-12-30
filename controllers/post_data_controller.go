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
		jsonErr := JsonResponseError{400, "wrong format"}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
	}

	email, found := dataRequest.Value["email"]
	if !found {
		fmt.Println("error param post")
		jsonErr := JsonResponseError{400, "missing param email"}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
	}

	name, found := dataRequest.Value["name"]
	if !found {
		fmt.Println("error param post")
		jsonErr := JsonResponseError{400, "missing param name"}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
	}

	//validate
	valid := validate.UserPostFormUpload{email[0], name[0]}
	err := valid.Validate()
	if err != nil {
		fmt.Println("error param post", MessageFrErrorValidate(err))
		jsonErr := JsonResponseError{400, MessageFrErrorValidate(err)}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
		return
	}

	poolJob.PushDataToQueue(payload.Payload{"test"})

	//response
	controller.Ctx.SetStatusCode(201)
	fmt.Fprintf(controller.Ctx, `{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
}
