package http_error_handle

import (
	"github.com/kataras/iris/v12"
	"handle-big-post-request/src/helper"
)

var LIST_STATUS_OKIE = []int{iris.StatusOK, iris.StatusCreated, iris.StatusAccepted}

func handleDefault(ctx iris.Context, app *iris.Application) {
	if helper.ItemExistsInSlice(LIST_STATUS_OKIE, ctx.GetStatusCode()) {
		return
	}

	app.Logger().Info("Request path: ", ctx.Path(), " ,Method: ", ctx.Method(), " ,StatusCode Not Success: ", ctx.GetStatusCode())
	return
}
