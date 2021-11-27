package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	_ "github.com/go-sql-driver/mysql"
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


	fmt.Println("after reset data buffeer, lenData", len(d.data))
	//todo  push to service save toDB
	d.insertBathToDB()
	// reset buffer data
	d.resetDataBufer()
	return true
}

func (d *DataSaveDb) resetDataBufer()  {
	d.data = make([]OneDataSaveDB, 0)
}

func (d *DataSaveDb) insertBathToDB() {
	db, err := sql.Open("mysql", "admin:1adphamnghia@tcp(127.0.0.1:3306)/handleBigPostRequest")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO post_data(name_user, email_user, detail_survey_user, created_at, updated_at) VALUES ('nghiapm', 'nghiaIt@gmail.com', 'test survey', 1, 2)"
	res, err := db.Exec(sql)

	if err != nil {
		fmt.Println("insert to DB error")
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		fmt.Println("get  last inserted row error")
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}
