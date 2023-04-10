package main

import (
	"github.com/ainmtsn1999/orm-book-api-deploy/app"
	"github.com/ainmtsn1999/orm-book-api-deploy/config"
)

func init() {
	// err := godotenv.Load()
	// if err != nil {
	// 	panic(err)
	// }

	err := config.InitGorm()
	if err != nil {
		panic(err)
	}
}

func main() {
	app.StartApplication()
}
