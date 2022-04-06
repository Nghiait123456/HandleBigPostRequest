package handle

import (
	"github.com/kataras/iris/v12"
	"handle-big-post-request/src/controllers"
	"handle-big-post-request/src/queue"
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
		booksAPI.Post("/", h.C.Create)
	}
}
