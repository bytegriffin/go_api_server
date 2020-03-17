package main

import (
	"go_api_server/dao"
	"go_api_server/routers"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	// dao.InitModel()
	r := routers.SetupRouter()
	r.Run()
}
