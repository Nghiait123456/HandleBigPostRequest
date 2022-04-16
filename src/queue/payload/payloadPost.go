package payload

import (
	"handle-big-post-request/src/models"
	"handle-big-post-request/src/queue/handle"
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
