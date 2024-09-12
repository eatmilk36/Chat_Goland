package Controller

import (
	"chat/Handler/Login"
	"chat/Repositories"
	"chat/Repositories/models"
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

	server.POST("Login", func(c *gin.Context) {
		database := Repositories.GormUserRepository{}.InitDatabase()

		// 初始化 UserRepository
		userRepo := models.NewGormUserRepository(database)

		// 注入到 LoginHandler
		loginHandler := Login.NewLoginHandler(userRepo)

		// 呼叫 業務邏輯
		loginHandler.LoginQueryHandler(c)
	})

	err := server.Run(":8080")
	if err != nil {
		panic("服務器啟動失敗")
	}
}
