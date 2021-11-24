package main

import (
	"sync"
	"time"
)

type Payload struct {
	test string
}

var muxtex = sync.Mutex{}

var dataSaveDb = DataSaveDb{[]OneDataSaveDB{}, 10000, muxtex}

func (p *Payload) Handle() bool {
	// fake request call other api verify
	for i := 0; i < 1000; i++ {

	}
	time.Sleep(300 * time.Millisecond)

	// build data :
	data := OneDataSaveDB{
		ModelPostSubmit{
			1,
			"nghiapm",
			"minhnghia.pham.it@gmail.com",
			"test khao sat thong tin user",
			time.Now().UTC().Unix(),
			time.Now().UTC().Unix(),
		},
	}

	dataSaveDb.addpendDataSaveDb(data)
	return true
}
