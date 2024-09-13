// @title Gin Swagger API Jeter

// @version 1.0

// @host localhost:8080

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @BasePath /
package main

import (
	"chat/Common"
	"chat/Controller"
	"chat/Middleware"
	_ "chat/docs"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	tokenString, err := Common.GenerateJWT("Jeter")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated JWT:", tokenString)

	// 驗證 JWT
	claims, err := Common.ValidateJWT(tokenString)
	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}

	// 成功驗證，輸出 Claims
	fmt.Printf("Token is valid! Username: %s\n", claims.Username)

	r := gin.Default()

	r.Use(Middleware.JWTAuthMiddleware())

	Controller.RouterInit(r)
}
