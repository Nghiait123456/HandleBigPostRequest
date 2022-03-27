package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"handle-big-post-request/src/config"
	"log"
	"os"
	"time"
)

var db *gorm.DB

// Init func load DB confirm then init DB connection
func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second * 2, // Slow SQL threshold
			LogLevel:      logger.Warn,     // Log level
			Colorful:      true,            // Disable color
		},
	)

	conf := config.GetAllConfig()
	username := conf.Database.UserName
	password := conf.Database.PassWord
	dbName := conf.Database.NameDB
	dbHost := conf.Database.Host
	dbPort := conf.Database.Port

	dbURI := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4", username, password, dbHost, dbPort, dbName)
	conn, err := gorm.Open(mysql.Open(dbURI), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	db = conn
}

func GetConnMysql() *gorm.DB {
	if db == nil {
		panic("Construct db after use")
	}

	return db
}
