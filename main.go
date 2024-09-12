// @title Gin Swagger API

// @host localhost:8080
// @BasePath /
package main

import (
	"chat/Controller"
	_ "chat/docs"
)

func main() {
	Controller.RouterInit()
}
