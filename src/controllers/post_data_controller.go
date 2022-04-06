package controllers

import (
	"github.com/kataras/iris/v12"
	"handle-big-post-request/src/logs_custom"
	"handle-big-post-request/src/queue"
	"handle-big-post-request/src/queue/payload"
	"handle-big-post-request/src/validate"
)

type PostDataController struct {
	PoolJob *queue.PoolJob
}

func (c *PostDataController) Create(ctx iris.Context) {
	prefixFc := "fc Create: "
	var formData validate.PostFormUpload
	err := ctx.ReadJSON(&formData)
	if err != nil {
		logs_custom.Logger().Warn(c.PreFixLog(prefixFc) + " please use content type application/json")
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{
			"status":  "false",
			"message": "please use content type application/json",
		})
		return
	}

	errVld := formData.Validate()
	if errVld != nil {
		logs_custom.Logger().Warn(c.PreFixLog(prefixFc) + " missing param required")
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{
			"status":  "false",
			"message": "missing param required",
		})
		return
	}

	c.PoolJob.PushDataToQueue(payload.Payload{formData.Name, formData.Email, formData.Detail})
	ctx.JSON(iris.Map{
		"status":  "success",
		"message": "done",
	})
}

func (c *PostDataController) PreFixLog(prefixFc string) string {
	return "class PostDataController : " + prefixFc
}
