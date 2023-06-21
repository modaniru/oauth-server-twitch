package main

import (
	"log"
	"time"

	"github.com/gojek/heimdall/httpclient"
	"github.com/modaniru/api-for-users/src/controller"
	"github.com/modaniru/api-for-users/src/repository"
	"github.com/modaniru/api-for-users/src/server"
	"github.com/modaniru/api-for-users/src/service"
	"github.com/modaniru/api-for-users/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	dsn := "host=localhost user=postgres password=qwerty dbname=postgres port=5556 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal(err.Error())
	}
	repository := repository.NewRepository(db)
	twitchRequest := utils.NewTwitchRequest(*httpclient.NewClient(
		httpclient.WithHTTPTimeout(10000 * time.Millisecond),
	))
	service := service.NewService(repository, twitchRequest)
	handler := controller.NewHandler(service)
	server := server.NewServer(handler.InitRouters(), "8080")
	server.Run()
}