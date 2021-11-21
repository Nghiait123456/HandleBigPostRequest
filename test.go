package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"sync"
)

func mainq() {
	router := gin.Default()

	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		//time.Sleep(time.Second * 100)
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.GET("/test", func(c *gin.Context) {
		var wg sync.WaitGroup
		for i := 0; i < 10000000; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				http.Get("http://localhost:8090/user/nghia")
				//fmt.Println("done one request %s", i)
			}(i)
		}
		wg.Wait()

		c.String(http.StatusOK, "successs")
	})



	fmt.Println("start run server")
	router.Run(":8090")
}

func main133() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()
	var wg sync.WaitGroup
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
		    http.Get(ts.URL)

			fmt.Println("done one request %s", i)
		}(i)
	}
	wg.Wait()
}