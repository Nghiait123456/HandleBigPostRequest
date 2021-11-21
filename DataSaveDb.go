package main

import "fmt"

type DataSaveDb struct {
	data [] OneDataSaveDB
	maxSize int
}

type OneDataSaveDB struct {
	oneData ModelPostSubmit
}

type ModelPostSubmit struct {
	id int
	nameUser string
	emailUser string
	detailSurveyUser string
	createdAt int64
	updateAt int64
}

func (d DataSaveDb) addpendDataSaveDb (data OneDataSaveDB, dSave *DataSaveDb)  {
	//append data
	if (len(dSave.data) <=  dSave.maxSize) {
		fmt.Print("dataSaveDB", dSave.data, "len", len(dSave.data))
		dSave.data = append(dSave.data, data)
	}

	fmt.Println("buffer max, push to service save DB")
	// push to service save toDB
}


