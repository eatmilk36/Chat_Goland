package Controller

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func RouterInit() {
	//database := Repositories.GormUserRepository{}.InitDatabase()

	server := gin.Default()

	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	//if mode := gin.Mode(); mode == gin.DebugMode {
	//	url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", "8080"))

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}

	server.GET("hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"hello": name})
	})

	server.GET("hello2/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"hello": name})
	})

	server.POST("/Login", UserController{}.GetUser)

	err := server.Run(":8080")
	if err != nil {
		panic("服務器啟動失敗")
	}
}
