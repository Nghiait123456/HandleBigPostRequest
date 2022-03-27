package http_error_handle

import "github.com/kataras/iris/v12"

func handleStatusInternalServerError(ctx iris.Context, app *iris.Application) {
	app.Logger().Warn("InternalServerError: Request path: ", ctx.Path(), " ,Method: ", ctx.Method(), " ,StatusCode Not Success: ", ctx.GetStatusCode())
	ctx.StatusCode(iris.StatusInternalServerError)
	ctx.JSON(iris.Map{
		"status":  "false",
		"message": "InternalServerError",
	})

	return
}
