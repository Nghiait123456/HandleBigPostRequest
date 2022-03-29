package w_pool_remote

import (
	"fmt"
	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/kataras/iris/v12"
	"handle-big-post-request/src/config"
	"handle-big-post-request/src/logs_custom"
	"io"
	"log"
	"os"
	"os/signal"
)

// Make a redis pool
var redisPool = &redis.Pool{
	MaxActive: 5,
	MaxIdle:   5,
	Wait:      true,
	Dial: func() (redis.Conn, error) {
		Dial := redis.DialPassword(config.GetAllConfig().Redis.PassWork)
		c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", config.GetAllConfig().Redis.Host, config.GetAllConfig().Redis.Port), Dial)
		if err != nil {
			panic(err.Error())
		}

		return c, err
	},
}

type Context struct {
	customerID int64
}

var app *iris.Application

func init() {
	fmt.Println("in fc Init")
	config.Init("../../config.yml")
	fmt.Println(config.GetAllConfig())
}

func main1() {
	// cf log
	f := logs_custom.NewLogFile("../log_file/")
	defer f.Close()
	app = iris.New()
	app.Logger().SetOutput(io.MultiWriter(f, os.Stdout))
	app.Logger().SetOutput(f)
	app.Logger().SetLevel("debug")
	app.Logger().Info("start worker_pool_remote")

	fmt.Println("in FC main")
	// Make a new pool. Arguments:
	// Context{} is a struct that will be the context for the request.
	// 10 is the max concurrency
	// "my_app_namespace" is the Redis namespace
	// redisPool is a Redis pool
	fmt.Println("start new worker Pool")
	pool := work.NewWorkerPool(Context{}, 10, "my_app_namespace", redisPool)

	fmt.Println("start middeware")
	// Add middleware that will be executed for each job
	pool.Middleware((*Context).Log)
	pool.Middleware((*Context).FindCustomer)

	// Map the name of jobs to handler functions
	pool.Job("send_email", (*Context).SendEmail)

	// Customize options:
	pool.JobWithOptions("export", work.JobOptions{Priority: 10, MaxFails: 1}, (*Context).Export)

	// Start processing jobs

	fmt.Println("start poll")
	pool.Start()
	fmt.Println("end worker")

	//add job to queue
	var enqueuer = work.NewEnqueuer("my_app_namespace", redisPool)
	// Enqueue a job named "send_email" with the specified parameters.
	_, err := enqueuer.Enqueue("send_email", work.Q{"address": "test@example.com", "subject": "hello world", "customer_id": 4})
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a signal to quit:
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)
	<-signalChan

	fmt.Println("stop worker ")
	// Stop the pool
	pool.Stop()
}

func (c *Context) Log(job *work.Job, next work.NextMiddlewareFunc) error {
	fmt.Println("Starting job: ", job.Name)
	return next()
}

func (c *Context) FindCustomer(job *work.Job, next work.NextMiddlewareFunc) error {
	// If there's a customer_id param, set it in the context for future middleware and handlers to use.
	if _, ok := job.Args["customer_id"]; ok {
		c.customerID = job.ArgInt64("customer_id")
		if err := job.ArgError(); err != nil {
			return err
		}
	}

	return next()
}

func (c *Context) SendEmail(job *work.Job) error {
	fmt.Println("in job sendEmail")
	// Extract arguments:
	addr := job.ArgString("address")
	subject := job.ArgString("subject")
	if err := job.ArgError(); err != nil {
		fmt.Printf("wrong format job")
		return err
	}
	fmt.Printf(" addr : %s, subject: %s \n", addr, subject)

	return nil
}

func (c *Context) Export(job *work.Job) error {
	return nil
}
