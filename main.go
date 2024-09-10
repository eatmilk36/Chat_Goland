package main

import (
	"chat/Handler/Login"
	"chat/Repositories"
	"chat/Repositories/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	Repositories.InitDatabase()

	id, err2 := models.GetUserByAccountAndPassword("Jeter", "MD5")
	if err2 != nil {
		return
	}

	fmt.Println("id:", id)

	server := gin.Default()

	server.GET("hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"hello": name})
	})

	server.GET("hello2/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"hello": name})
	})

	server.POST("Login", Login.LoginHandler)

	err := server.Run(":8080")
	if err != nil {
		panic("服務器啟動失敗")
	}
}
