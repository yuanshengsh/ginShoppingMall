package service

import (
	"ginShoppingMall/app/model"
	boot "ginShoppingMall/bootstrap"
)

var user model.User

func GetUser() interface{} {
	boot.DB.Find(&user)
	return user
}
