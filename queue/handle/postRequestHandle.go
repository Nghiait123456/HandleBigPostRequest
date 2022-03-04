package handle

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"handle-big-post-request/config"
	"handle-big-post-request/connect/database"
	"handle-big-post-request/models"
	"handle-big-post-request/repository"
	"sync"
	"time"
)

type DataSaveDb struct {
	Data []models.PostSubmit
	// calculation for Max placeHolder DB and Table Insert
	MaxSize int
	Mu      sync.Mutex
}

func (d *DataSaveDb) AddpendDataSaveDb(data models.PostSubmit) bool {
	// if multil routine use func, if sync ==> use mutex
	// todo update use atomic ==> up performance
	d.Mu.Lock()
	defer d.Mu.Unlock()
	//append data
	if len(d.Data) < d.MaxSize {
		d.Data = append(d.Data, data)
		return true
	}

	//todo  push to service save toDB
	//todo cmt for test aws without DB
	//ok := d.insertBathToDB()
	//if ok != nil {
	//	fmt.Println("insert Batch error")
	//	return false
	//}

	fmt.Println("insert Batch success")
	// reset buffer data
	d.resetDataBufer()
	return true
}

func (d *DataSaveDb) resetDataBufer() {
	d.Data = make([]models.PostSubmit, 0)
}

func (d *DataSaveDb) insertBathToDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	R := repository.NewPostDataRepository(database.GetConnect(config.AllConfig.Database.Type))
	return R.BatchInsert(ctx, &d.Data)
}
