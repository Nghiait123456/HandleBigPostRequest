package repo

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func connectToDBMysql(typeDB string) (db *gorm.DB, err error) {
	switch typeDB {
	case "mysql":
		dsn := "admin:1adphamnghia@tcp(172.17.0.2:3306)/handleBigPostRequest?charset=utf8mb4&parseTime=True&loc=Local"
		return gorm.Open(mysql.Open(dsn), &gorm.Config{})
	default:
		log.Println("not support typeDB", typeDB)
		return nil, errors.New("support typeDB")
	}
}

func InsertBatch(data interface{}) bool {
	db, err := connectToDBMysql("mysql")
	if err != nil {
		log.Println(err)
		return false
	}

	db.Create(data)
	return true
}
