package Login

import (
	"chat/Repositories"
	"chat/Repositories/models"
	"chat/Test"
	"github.com/go-playground/assert/v2"
	"testing"
	"time"
)

func TestLoginHandler(t *testing.T) {
	result := 2 + 3
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestUserInsert(t *testing.T) {
	// 使用共用的 SetupTestDB 和 ResetDB
	db := Test.SetupTestDB(t)
	Test.ResetDB(db)

	Repositories.DB = db

	// 插入測試資料
	user := models.User{
		Account:     "Jeter",
		Password:    "MD5",
		Id:          1,
		CreatedTime: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}

	user = models.User{
		Account:     "Jeter2",
		Password:    "MD5",
		Id:          2,
		CreatedTime: time.Now(),
	}

	if err := db.Create(&user).Error; err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}

	userModel, _ := models.GetUserByAccountAndPassword("Jeter", "MD5")

	assert.Equal(t, "Jeter", userModel.Account)
}
