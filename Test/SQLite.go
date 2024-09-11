package Test

import (
	"chat/Repositories/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
	"time"
)

// 初始化資料庫，使用 SQLite 內存模式
func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自動遷移資料庫結構
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestGormWithSQLite(t *testing.T) {
	// 初始化測試資料庫
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	// 插入測試資料
	var testUser = models.User{
		Account:     "Jeter",
		Password:    "Sara",
		Id:          1,
		CreatedTime: time.Now(),
		//CreatedTime: time.Date(2024, time.September, 11, 10, 30, 0, 0, time.UTC), // 設置為 2024-09-11 10:30:00 UTC
	}
	if err := db.Create(&testUser).Error; err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}

	// 查詢資料
	var user models.User
	if err := db.First(&user, 1).Error; err != nil {
		t.Fatalf("Failed to query test user: %v", err)
	}

	// 驗證結果
	if user.Account != "Test User" {
		t.Errorf("Expected user name to be 'Test User', got '%s'", user.Account)
	}
}
