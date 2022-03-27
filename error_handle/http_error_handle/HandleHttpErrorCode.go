package http_error_handle

import (
	"github.com/kataras/iris/v12"
	"handle-big-post-request/config"
)

type HandleHttp struct {
	App *iris.Application
}

func (hdl *HandleHttp) ResignHandleHttpError(ctx iris.Context) {
	config := config.GetAllConfig()
	if config.Mode.ModeRun != "debug" {
		return
	}

	switch ctx.GetStatusCode() {
	case iris.StatusInternalServerError:
		handleStatusInternalServerError(ctx, hdl.App)
	default:
		handleDefault(ctx, hdl.App)
	}
}

func NewHandleHttpError(app *iris.Application) *HandleHttp {
	return &HandleHttp{app}
}
