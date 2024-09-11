package Controller

import (
	"chat/Handler/Login"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterInit() {

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
