package models

import (
	"time"
)

type User struct {
	//gorm.Model
	Account     string    `json:"Account"`
	Password    string    `json:"Password"`
	Id          int       `json:"Id"`
	CreatedTime time.Time `json:"CreatedTime"`
}
