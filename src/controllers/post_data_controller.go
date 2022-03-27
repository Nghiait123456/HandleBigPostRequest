package controllers

import (
	"github.com/kataras/iris/v12"
	"handle-big-post-request/src/queue"
	"handle-big-post-request/src/queue/payload"
)

type PostDataController struct {
	PoolJob *queue.PoolJob
}

type ResponseSuccess struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (c *PostDataController) Create(Ctx iris.Context) {
	email := Ctx.PostValueDefault("email", "")
	if "" == email {
		Ctx.StatusCode(iris.StatusBadRequest)
		Ctx.JSON(iris.Map{
			"status":  "false",
			"message": "email is require",
		})
		return
	}

	name := Ctx.PostValueDefault("name", "")
	if "" == name {
		Ctx.StatusCode(iris.StatusBadRequest)
		Ctx.JSON(iris.Map{
			"status":  "false",
			"message": "name is require",
		})
		return
	}

	detail := Ctx.PostValueDefault("detail", "detail default")
	c.PoolJob.PushDataToQueue(payload.Payload{name, email, detail})

	Ctx.JSON(iris.Map{
		"status":  "success",
		"message": "done",
	})
}
