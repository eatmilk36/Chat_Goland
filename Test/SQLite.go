package Test

import (
	"chat/Repositories/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
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

// TestMain 用來在所有測試之前初始化共用資源，並在測試完成後釋放資源
func TestMain(m *testing.M) {
	// 初始化資料庫
	db := SetupTestDB(nil)
	ResetDB(db)
	//database := Repositories.GormUserRepository{}.InitDatabase() // 把全域的DB改成SQLite

	//database = db

	// 建立 GormUserRepository 的實例並初始化資料庫
	_ = models.NewGormUserRepository(db) // 假設有 NewGormUserRepository 函式來初始化

	// 執行測試
	sqlDB, err := db.DB() // 取得底層的 sql.DB
	if err == nil {
		sqlDB.Close() // 釋放資料庫資源
	}

	// 執行測試並取得結果狀態碼
	code := m.Run()

	// 結束測試
	os.Exit(code)
}
