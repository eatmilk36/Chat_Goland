// @title Gin Swagger API
// @version 1.0
// @description This is a sample Gin server.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url http://example.com/license

// @host localhost:8080
// @BasePath /
package main

import (
	"chat/Controller"
	_ "chat/docs"
)

func main() {

	//docs.SwaggerInfo.BasePath = "/"
	//ctx := context.Background()
	//
	//// 使用載入的配置初始化 Redis 客戶端
	//redisService := Redis.NewRedisService()
	//
	//// 示例：設置 Redis 的 Hash 值
	//err := redisService.HashSet(ctx, "MyHash", "field1", "value1")
	//if err != nil {
	//	log.Fatalf("failed to set hash: %v", err)
	//}
	//
	//// 示例：獲取 Redis 的 Hash 值
	//value, err := redisService.HashGet(ctx, "MyHash", "field1")
	//if err != nil {
	//	log.Fatalf("failed to get hash: %v", err)
	//}
	//
	//log.Printf("Value from Hash: %s", value)

	Controller.RouterInit()
}
