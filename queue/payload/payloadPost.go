package payload

import (
	"handle-big-post-request/models/rawQuery"
	"sync"
	"time"
)

type Payload struct {
	Name   string
	Email  string
	Detail string
}

var muxtex = sync.Mutex{}

var dataSaveDb = rawQuery.DataSaveDb{[]rawQuery.PostSubmit{}, 10000, muxtex}

func (p *Payload) Handle() bool {
	// fake request call other api verify
	for i := 0; i < 1000; i++ {

	}
	time.Sleep(300 * time.Millisecond)

	// build data :
	data := rawQuery.PostSubmit{
		p.Name,
		p.Email,
		p.Detail,
		time.Now().UTC().Unix(),
		time.Now().UTC().Unix(),
	}

	dataSaveDb.AddpendDataSaveDb(data)
	return true
}
