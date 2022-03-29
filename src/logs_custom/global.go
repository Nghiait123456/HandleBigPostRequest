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

	body, _ := ctx.GetBody()
	if g.IsPrintGlobal == true {
		fmt.Println("Request path: ", ctx.Path(), ", request Body: ", body)
	}

	if g.IsPrintGlobal == true {
		ctx.Application().Logger().Info("Request path: ", ctx.Path(), ", request Body: ", body)
	}
}

func (g *Global) globalMiddlewareLogs(ctx iris.Context) {
	g.logFullRequest(ctx)
	ctx.Next()
}

func (g *Global) ResignGlobalLog(app *iris.Application) {
	app.UseGlobal(g.globalMiddlewareLogs)
}
