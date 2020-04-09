package controller

import (
	"github.com/gin-gonic/gin"
	"go_api_server/dao"
	"go_api_server/models"
	"net/http"
	"strconv"
)

func CreateUser(c *gin.Context) {
	login_name := c.PostForm("login_name")
	if login_name == "" {
		c.JSON(http.StatusOK, gin.H{"error": "login_name is null."})
		return
	}
	user := new(models.User)
	user.Login_Name = login_name
	c.BindJSON(&user)

	err := dao.InsertUser(*user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetOneUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "FAILD ID."})
		return
	}
	user, err := dao.GetOneUser(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, user)
	}
}

func GetUserList(c *gin.Context) {
	userList, err := dao.GetUserList()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, userList)
	}
}

func UpdateOneUser(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "failed id."})
		return
	}
	login_name := c.PostForm("login_name")
	user := new(models.User)
	user.ID, _ = strconv.Atoi(id)
	user.Login_Name = login_name
	c.BindJSON(user)
	if err := dao.UpdateOneUser(*user); err != nil {
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
