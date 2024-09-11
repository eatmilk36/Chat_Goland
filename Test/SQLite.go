package Test

import (
	"chat/Repositories/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// SetupTestDB 共用的設置資料庫的函式
func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	// 自動遷移
	db.AutoMigrate(&models.User{})
	//db.AutoMigrate(&User{}, &Order{})

	return db
}

func ResetDB(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
	db.AutoMigrate(&models.User{})
	//db.Migrator().DropTable(&User{}, &Order{})
	//db.AutoMigrate(&User{}, &Order{})
}
