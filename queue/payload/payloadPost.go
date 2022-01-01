package payload

import (
	"handle-big-post-request/queue/handle"
	"handle-big-post-request/repo"
	"sync"
	"time"
)

type Payload struct {
	Name   string
	Email  string
	Detail string
}

var muxtex = sync.Mutex{}

var dataSaveDb = handle.DataSaveDb{[]repo.PostSubmit{}, 10000, muxtex}

func (p *Payload) Handle() bool {
	// fake request call other api verify
	for i := 0; i < 1000; i++ {

	}
	time.Sleep(300 * time.Millisecond)

	// build data :
	data := repo.PostSubmit{
		p.Name,
		p.Email,
		p.Detail,
		time.Now().UTC().Unix(),
		time.Now().UTC().Unix(),
	}

	dataSaveDb.AddpendDataSaveDb(data)
	return true
}
