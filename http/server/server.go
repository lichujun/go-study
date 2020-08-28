package main

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
	"time"
)

func helloWorld(request *restful.Request, response *restful.Response) {
	fmt.Fprintf(response, "hello world")
}

func main() {
	// container 同样实现了http.Hadler接口
	container := restful.NewContainer()
	webService := new(restful.WebService)
	webService.Route(webService.GET("/hello").To(helloWorld))
	container.Add(webService)
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      container,
		Addr:         ":8080",
	}
	log.Fatal(server.ListenAndServe())
	/*http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("start server err" + err.Error())
	}*/
}
