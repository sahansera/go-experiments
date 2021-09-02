package model

import (
	"gintest/config"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (b *User) TableName() string {
	return "user"
}

func GetAllUsers(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}
