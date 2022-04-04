package logs_custom

import (
	"fmt"
	"github.com/kataras/iris/v12"
)

type Global struct {
	IsLogGlobal        bool
	IsPrintGlobal      bool
	IsUseLogGlobalMode bool
}

func (g *Global) logFullRequest(ctx iris.Context) {
	if !g.IsUseLogGlobalMode {
		return
	}

	if g.IsPrintGlobal == true {
		fmt.Println("Request path: %s ", ctx.Path())
	}

	if g.IsPrintGlobal == true {
		ctx.Application().Logger().Info("Request path: %s \n", ctx.Path())
	}
}

func (g *Global) globalMiddlewareLogs(ctx iris.Context) {
	g.logFullRequest(ctx)
	ctx.Next()
}

func (g *Global) ResignGlobalLog(app *iris.Application) {
	app.UseGlobal(g.globalMiddlewareLogs)
}
