package main

import (
	"github.com/emicklei/go-restful"
	"go-study/http/server/controller"
	"log"
	"net/http"
	"time"
)

func main() {
	// container 同样实现了http.Hadler接口t
	container := restful.NewContainer()
	services := controller.GetAllService()
	for _, service := range services {
		container.Add(service)
	}
	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      container,
		Addr:         ":9000",
	}
	log.Fatal(server.ListenAndServe())
	/*http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("start server err" + err.Error())
	}*/
}
