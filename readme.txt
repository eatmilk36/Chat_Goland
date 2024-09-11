	md5String := Commond.Md5Hash("hahawhy")

	// 輸出結果
	fmt.Println("MD5 雜湊值:", md5String)

	ctx := context.Background()

	// 初始化 Redis 客戶端
	client := Redis.NewRedisClient()
	redisService := Redis.NewRedisService(client)

	err := redisService.SetValue(ctx, "sandy", "sara")
	if err != nil {
		log.Fatalf("failed to set key: %v", err)
	}

	// 使用 Redis 服務獲取值
	value, err := redisService.GetValue(ctx, "sandy")
	if err != nil {
		log.Fatalf("failed to get key: %v", err)
	}

	fmt.Println("Value from Redis:", value)

	err = redisService.ListPush(ctx, "mylist", "sandy", "sara", "lin")
	if err != nil {
		log.Fatalf("failed to push to list: %v", err)
	}

	// 獲取 List 中的數值
	values2, err2 := redisService.ListRange(ctx, "mylist", 0, -1) // -1 表示到 List 的最後一個元素
	if err2 != nil {
		log.Fatalf("failed to get list range: %v", err2)
	}

	fmt.Println("ListRange:", values2)

	err3 := redisService.HashSet(ctx, "cat", "apple2", "meow2")
	if err3 != nil {
		log.Fatalf("failed to set hash: %v", err3)
	}

	// 獲取 Hash 中的字段值
	value3, err3 := redisService.HashGet(ctx, "cat", "apple2")
	if err3 != nil {
		log.Fatalf("failed to get hash: %v", err3)
	}

	log.Printf("Value from Hash: %s", value3)

	Repositories.InitDatabase()

	id, err2 := models.GetUserByAccountAndPassword("Jeter", "MD5")
	if err2 != nil {
		return
	}

	fmt.Println("id:", id)