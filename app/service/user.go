package service

import (
	"ginShoppingMall/app/model"
	boot "ginShoppingMall/bootstrap"
	"strconv"
	"time"
)

var u model.User

//func GetUser() interface{} {
//	//boot.DB.Find(&user)
//	//return user
//}

func AddUser(user *model.User) (err error) {
	boot.DB.Create(&model.User{
		Mobile:   user.Mobile,
		Name:     user.Name,
		Password: user.Password,
	})
	return nil
}

func CheckUserExistByMobileOrName(mobile string, name string) bool {
	boot.DB.Where("mobile = ?", "mobile").Or("name = ?", "name").Find(&u)
	//if user {
	//	return true
	//}
	return false
}

func ExpireAt() string {
	currentTime := time.Now()
	m, _ := time.ParseDuration("+" + strconv.Itoa(boot.Config.Token.JwtTokenCreatedExpireAt) + "m")
	return currentTime.Add(m).Format("2006-01-02 15:04:05")
}
