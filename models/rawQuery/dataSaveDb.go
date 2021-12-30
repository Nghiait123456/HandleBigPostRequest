package rawQuery

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"sync"
)

type DataSaveDb struct {
	Data []PostSubmit
	// calculation for Max placeHolder DB and Table Insert
	MaxSize int
	Mu      sync.Mutex
}

type PostSubmit struct {
	NameUser         string
	EmailUser        string
	DetailSurveyUser string
	CreatedAt        int64
	UpdatedAt        int64
}

func (d *DataSaveDb) AddpendDataSaveDb(data PostSubmit) bool {
	// if multil routine use func, if sync ==> use mutex
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
	d.Data = make([]PostSubmit, 0)
}

func (d *DataSaveDb) insertBathToDB() bool {
	dsn := "admin:1adphamnghia@tcp(127.0.0.1:3306)/handleBigPostRequest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return false
	}

	db.Create(&d.Data)
	return true
}
