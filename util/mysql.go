package util

import (
	"fmt"
	boot "ginShoppingMall/bootstrap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func StartMysql() *gorm.DB {
	// 数据库配置
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		boot.Config.MySQL.User,
		boot.Config.MySQL.Pass,
		boot.Config.MySQL.Host,
		boot.Config.MySQL.Port,
		boot.Config.MySQL.DataBase,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	return db
}
