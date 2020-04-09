package routers

import (
	"github.com/gin-gonic/gin"
	"go_api_server/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/user", controller.CreateUser)
		v1Group.GET("/user/:id", controller.GetOneUser)
		v1Group.GET("/user", controller.GetUserList)
		v1Group.PUT("/user/:id", controller.UpdateOneUser)
		v1Group.DELETE("/user/:id", controller.DeleteOneUser)
	}

	return r
}
