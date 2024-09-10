package models

import (
	"time"
)

type User struct {
	//gorm.Model
	Account     string
	Password    string
	Id          int
	CreatedTime time.Time
}
