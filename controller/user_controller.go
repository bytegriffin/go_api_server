package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_api_server/dao"
	"go_api_server/models"
	"net/http"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)

	err := dao.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUserList(c *gin.Context) {
	userList, err := dao.GetUserList()
	fmt.Println(userList)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, userList)
	}
}

func UpdateOneUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "FAILD ID."})
		return
	}
	user, err := dao.GetOneUser(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.BindJSON(*user)
	if err = dao.UpdateOneUser(user); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func DeleteOneUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "FAILD ID."})
		return
	}
	if err := dao.DeleteOneUser(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
