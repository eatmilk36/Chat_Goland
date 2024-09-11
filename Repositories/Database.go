package Repositories

import (
	"chat/Config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

func InitDatabase() {
	config, err2 := Config.LoadConfig()
	if err2 != nil {
		panic("Config file load failed")
	}
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True",
		config.MySql.UserName, config.MySql.Password, config.MySql.Address, config.MySql.Port, config.MySql.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB, err := DB.DB()

	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return
	}

	DB.SetConnMaxLifetime(time.Duration(config.MySql.MaxLifetime) * time.Second)
	DB.SetMaxIdleConns(config.MySql.MaxIdleConnections)
	DB.SetMaxOpenConns(config.MySql.MaxOpenConnections)
}
