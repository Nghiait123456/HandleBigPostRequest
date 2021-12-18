package rawQuery

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sync"
)

type DataSaveDb struct {
	Data []OneDataSaveDB
	// calculation for Max placeHolder DB and Table Insert
	MaxSize int
	Mu      sync.Mutex
}

type OneDataSaveDB struct {
	OneData ModelPostSubmit
}

type ModelPostSubmit struct {
	Id               int
	NameUser         string
	EmailUser        string
	DetailSurveyUser string
	CreatedAt        int64
	UpdateAt         int64
}

func (d *DataSaveDb) AddpendDataSaveDb(data OneDataSaveDB) bool {
	// if multil routine use func, if sync
	// use mutex
	// todo update use atomic ==> up performance
	d.Mu.Lock()
	defer d.Mu.Unlock()
	//append data
	if len(d.Data) < d.MaxSize {
		d.Data = append(d.Data, data)
		fmt.Println("dataSaveDB", d.Data, "len", len(d.Data), "maxSize", d.MaxSize)
		return true
	}

	fmt.Println("save data to DB, lenData", len(d.Data))

	fmt.Println("after reset data buffeer, lenData", len(d.Data))
	//todo  push to service save toDB
	d.insertBathToDB()
	// reset buffer data
	d.resetDataBufer()
	return true
}

func (d *DataSaveDb) resetDataBufer() {
	d.Data = make([]OneDataSaveDB, 0)
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
