package models

import "chat/Repositories"

func CreateUser(user *User) error {
	return Repositories.DB.Create(user).Error
}

func GetUserByID(id uint) (*User, error) {
	var user User
	err := Repositories.DB.First(&user, id).Error
	return &user, err
}

func UpdateUser(user *User) error {
	return Repositories.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	return Repositories.DB.Delete(&User{}, id).Error
}

func GetUserByAccountAndPassword(account string, password string) (*User, error) {
	var user User
	if err := Repositories.DB.Where("account = ? AND password = ?", account, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
