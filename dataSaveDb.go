package main

import "fmt"

type DataSaveDb struct {
	data    []OneDataSaveDB
	maxSize int
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
	//append data
	if len(d.data) <= d.maxSize {
		fmt.Print("dataSaveDB", d.data, "len", len(d.data))
		d.data = append(d.data, data)
	}

	fmt.Println("buffer max, push to service save DB")
	//todo  push to service save toDB
}
