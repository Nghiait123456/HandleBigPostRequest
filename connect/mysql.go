package connect

import "fmt"

import (
_ "github.com/go-sql-driver/mysql"
"github.com/jinzhu/gorm"
)
// SetupDB : initializing mysql database
func SetupDB() *gorm.DB {
	USER := "admin"
	PASS := "1adphamnghia"
	HOST := "localhost"
	PORT := "3306"
	DBNAME := "handleBigPostRequest"
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}
