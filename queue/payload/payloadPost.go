package payload

import (
	"github.com/valyala/fasthttp"
	"handle-big-post-request/models"
	"handle-big-post-request/queue/handle"
	"sync"
	"time"
)

type Payload struct {
	Name   string
	Email  string
	Detail string
}

var muxtex = sync.Mutex{}

var dataSaveDb = handle.DataSaveDb{
	[]models.PostSubmit{},
	10000,
	muxtex,
}

func (p *Payload) Handle() bool {
	//fake call api post Data
	client := fasthttp.Client{}
	var get []byte
	client.GetTimeout(get, "https://www.cloudflare.com/", 150*time.Millisecond)

	// build data :
	data := models.PostSubmit{
		p.Name,
		p.Email,
		p.Detail,
		time.Now().UTC().Unix(),
		time.Now().UTC().Unix(),
	}

	return dataSaveDb.AddpendDataSaveDb(data)
}
