package controller

import "github.com/emicklei/go-restful/v3"

var services []*restful.WebService

func RegisterWebService(service *restful.WebService) {
	services = append(services, service)
}

func GetAllService() []*restful.WebService {
	return services
}
