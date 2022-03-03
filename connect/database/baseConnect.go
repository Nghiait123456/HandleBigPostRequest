package database

import (
	"gorm.io/gorm"
)

func GetConnect(typeDB string) *gorm.DB {
	var db *gorm.DB
	switch typeDB {
	case "mysql":
		db = GetConnMysql()
	default:
		panic("UnKnow DB config")
	}

	return db
}
