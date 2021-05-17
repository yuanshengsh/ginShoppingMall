package model

import (
	boot "ginShoppingMall/bootstrap"
	"time"
)

type User struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Mobile string `json:"mobile"`
	//MobileVerifiedAt string `json:"mobile_verified_at"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SaveUser(u *User) error {
	if err := boot.DB.Create(u).Error; err != nil {
		return err
	}
	return nil
}

func FindUserByMobile(mobile string, u *User) error {
	if err := boot.DB.Where("mobile = ?", mobile).First(&u).Error; err != nil {
		return err
	}
	return nil

}

func FindUserLogin(name string, password string, u *User) error {
	if err := boot.DB.Where("name = ? and password = ?", name, password).First(&u).Error; err != nil {
		return err
	}
	return nil

}

func (User) TableName() string {
	return "user"
}
