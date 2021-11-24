package main

import (
	"time"
)

type Payload struct {
	test string
}

var dataSaveDb = DataSaveDb{[]OneDataSaveDB{}, 10000}

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

	//dataSaveDb.data = append(dataSaveDb.data, data)
	//fmt.Print("dataSave", dataSaveDb.data, "len", len(dataSaveDb.data))
	dataSaveDb.addpendDataSaveDb(data)
	return true
}
