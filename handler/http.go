package handler

import "github.com/ainmtsn1999/orm-book-api-deploy/service"

type HttpServer struct {
	app service.ServiceInterface
}

func NewHttpServer(app service.ServiceInterface) HttpServer {
	return HttpServer{app: app}
}
