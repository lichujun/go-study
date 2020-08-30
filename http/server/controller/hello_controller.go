package controller

import (
	"fmt"
	"github.com/emicklei/go-restful"
)

func init() {
	webService := new(restful.WebService)
	webService.Route(webService.GET("/").To(hello))
	RegisterWebService(webService)
}

func hello(_ *restful.Request, response *restful.Response) {
	fmt.Fprintf(response, "hello world")
}
