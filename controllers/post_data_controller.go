package controllers

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"handle-big-post-request/queue"
	"handle-big-post-request/queue/payload"
	"handle-big-post-request/validate"
)

type PostDataController struct {
	//PostData models.PostData
	//Sessions validate.UserPostFormUpload
	Ctx *fasthttp.RequestCtx
}

type ResponseSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (controller *PostDataController) Create(poolJob queue.PoolJob) {
	// get data request
	dataRequest, ok := controller.Ctx.MultipartForm()
	if ok != nil {
		fmt.Println("error param post")
		jsonErr := JsonResponseError{400, "wrong format"}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
		return
	}

	email, found := dataRequest.Value["email"]
	if !found {
		fmt.Println("error param post")
		jsonErr := JsonResponseError{400, "missing param email"}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
		return
	}

	fmt.Println("Email", email, len(email))

	name, found := dataRequest.Value["name"]
	if !found {
		fmt.Println("error param post")
		jsonErr := JsonResponseError{400, "missing param name"}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
		return
	}

	detail, found := dataRequest.Value["detail"]
	if !found {
		fmt.Println("error param post")
		jsonErr := JsonResponseError{400, "missing param detail"}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
		return
	}

	//validate
	valid := validate.UserPostFormUpload{email[0], name[0], detail[0]}
	err := valid.Validate()
	if err != nil {
		fmt.Println("error param post", MessageFrErrorValidate(err))
		jsonErr := JsonResponseError{400, MessageFrErrorValidate(err)}
		rs := ResponseError{controller.Ctx, 400, jsonErr}
		rs.ResponseError()
		return
	}

	poolJob.PushDataToQueue(payload.Payload{name[0], email[0], detail[0]})

	//response success
	rsSuccess := ResponseSuccess{true, "success"}
	ResponseJson(controller.Ctx, rsSuccess, 201)
}
