package main

import (
	"log"
	"time"
)

const MAX_QUEUE = 1000

type Payload struct {
	test string
}

var dataSaveDb = DataSaveDb{[]OneDataSaveDB{}, 10000}

func (p *Payload) Handle() bool {
	// fake request post to s3
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
	dataSaveDb.addpendDataSaveDb(data, &dataSaveDb)
	return true
}

func handleJob(p *Payload) bool {
	go p.Handle()
	return true
}

var Queue chan Payload

func initQueue() {
	Queue = make(chan Payload, MAX_QUEUE)
}

func payloadHandler(payload Payload) {
	// push to queue
	log.Println("start push to queue")
	Queue <- payload
	log.Println("end push to queue")
}
