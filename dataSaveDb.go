package main

import (
	"fmt"
	"sync"
)

type DataSaveDb struct {
	data    []OneDataSaveDB
	// calculation for Max placeHolder DB and Table Insert
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

func (d *DataSaveDb) addpendDataSaveDb(data OneDataSaveDB) bool {
	// if multil routine use func, if sync
	// use mutex
	// todo update use atomic ==> up performance
	d.mu.Lock()
	defer d.mu.Unlock()
	//append data
	if len(d.data) < d.maxSize {
		d.data = append(d.data, data)
		//fmt.Println("dataSaveDB", d.data, "len", len(d.data), "maxSize", d.maxSize)
		return true
	}

	fmt.Println("save data to DB, lenData", len(d.data))

	// reset buffer data
	d.resetDataBufer()

	fmt.Println("after reset data buffeer, lenData", len(d.data))
	//todo  push to service save toDB
	return true
}

func (d *DataSaveDb) resetDataBufer()  {
	d.data = make([]OneDataSaveDB, 0)
}
