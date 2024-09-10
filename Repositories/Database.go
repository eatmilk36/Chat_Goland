package Repositories

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

const (
	UserName           string = "root"
	Password           string = "jeter"
	Addr               string = "127.0.0.1"
	Port               int    = 3306
	Database           string = "chat"
	MaxLifetime        int    = 10
	MaxOpenConnections int    = 10
	MaxIdleConnections int    = 10
)

func InitDatabase() {
	addr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", UserName, Password, Addr, Port, Database)
	var err error
	DB, err = gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	//
	//DB, err := DB.DB()
	//
	//if err != nil {
	//	fmt.Println("Failed to connect to database:", err)
	//	return
	//}
	//
	//DB.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	//DB.SetMaxIdleConns(MaxIdleConnections)
	//DB.SetMaxOpenConns(MaxOpenConnections)
}
