package main

import (
	"chat/Controller"
	"chat/Redis"
	"chat/Repositories"
	"chat/Repositories/models"
	"context"
	"fmt"
	"log"
)

func main() {

	ctx := context.Background()

	// 使用載入的配置初始化 Redis 客戶端
	redisService := Redis.NewRedisService()

	// 示例：設置 Redis 的 Hash 值
	err := redisService.HashSet(ctx, "MyHash", "field1", "value1")
	if err != nil {
		log.Fatalf("failed to set hash: %v", err)
	}

	// 示例：獲取 Redis 的 Hash 值
	value, err := redisService.HashGet(ctx, "MyHash", "field1")
	if err != nil {
		log.Fatalf("failed to get hash: %v", err)
	}

	log.Printf("Value from Hash: %s", value)

	Repositories.InitDatabase()

	user, err := models.GetUserByAccountAndPassword("jeter", "MD5")
	if err != nil {
		fmt.Println("sql query failed", err)
	}

	fmt.Println("result:", user)

	Controller.RouterInit()
}
