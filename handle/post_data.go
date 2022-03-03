package handle

import (
	"github.com/kataras/iris/v12"
	"handle-big-post-request/controllers"
	"handle-big-post-request/queue"
)

type PostData struct {
	App              *iris.Application
	PoolWorkerUpload *queue.PoolJob
	C                *controllers.PostDataController
}

func ResignRoutePostData(h *PostData) {
	booksAPI := h.App.Party("/uploadForm")
	{
		// POST: http://localhost:8080/uploadForm
		booksAPI.Post("/", h.CreatePostDataControler)
	}
}

func (h *PostData) CreatePostDataControler(ctx iris.Context) {
	h.C.Create(ctx)
}
