package handle

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/valyala/fasthttp"
	"handle-big-post-request/src/config"
	"handle-big-post-request/src/connect/database"
	"handle-big-post-request/src/models"
	"handle-big-post-request/src/repository"
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

	// push data to DB or push to another service to save
	// anything ==> target batch Save Data, batch insert but not save every-Pne

	// if user save DB
	//ok := d.insertBathToDB()
	//if ok != nil {
	//	fmt.Println("insert Batch error")
	//	return false
	//}

	//fake call api post Data to service save data
	client := fasthttp.Client{}
	var get []byte
	client.GetTimeout(get, "https://www.cloudflare.com/", 150*time.Millisecond)

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
