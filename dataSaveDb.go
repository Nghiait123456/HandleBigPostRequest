package main

import (
	"fmt"
	"sync"
)

type DataSaveDb struct {
	data    []OneDataSaveDB
	maxSize int
	mu       sync.Mutex
}

type OneDataSaveDB struct {
	oneData ModelPostSubmit
}

type ModelPostSubmit struct {
	id               int
	nameUser         string
	emailUser        string
	detailSurveyUser string
	createdAt        int64
	updateAt         int64
}

func (d *DataSaveDb) addpendDataSaveDb(data OneDataSaveDB) {
	// if multil routine use func, if sync
	// use mutex
	// todo update use atomic ==> up performance
	d.mu.Lock()
	defer d.mu.Unlock()
	//append data
	if len(d.data) <= d.maxSize {
		fmt.Print("dataSaveDB", d.data, "len", len(d.data))
		d.data = append(d.data, data)
	}

	fmt.Println("buffer max, push to service save DB")
	//todo  push to service save toDB
}
