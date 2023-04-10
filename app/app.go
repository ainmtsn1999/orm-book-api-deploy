package app

import (
	"fmt"
	"os"

	"github.com/ainmtsn1999/orm-book-api-deploy/config"
	"github.com/ainmtsn1999/orm-book-api-deploy/handler"
	"github.com/ainmtsn1999/orm-book-api-deploy/repository"
	"github.com/ainmtsn1999/orm-book-api-deploy/route"
	"github.com/ainmtsn1999/orm-book-api-deploy/service"
	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	service := service.NewService(repo)
	server := handler.NewHttpServer(service)

	route.RegisterApi(router, server)

	port := os.Getenv("PORT")
	router.Run(fmt.Sprintf(":%s", port))
}
