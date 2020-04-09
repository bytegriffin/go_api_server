package main

import (
	"go_api_server/config"
	"go_api_server/dao"
	"go_api_server/routers"
	"log"
)

func init() {
	config.Setup()
}

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	r := routers.SetupRouter()
	log.Print("http server %s starting...", config.AppConfig.AppName)
	r.Run()
}
