package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

func StartMysql(MyEnv map[string]string) *gorm.DB {
//	// 数据库配置
//	SetDatabasePrefix(MyEnv["DB_PREFIX"])
//	dsn := MyEnv["DB_USERNAME"] + ":" + MyEnv["DB_PASSWORD"] + "@(" + MyEnv["DB_HOST"] + ":" + MyEnv["DB_PORT"] + ")/" + MyEnv["DB_DATABASE"] + "?parseTime=true&loc=Asia%2FShanghai&charset=" + MyEnv["DB_CHARSET"]
//	if MyEnv["DB_DSN"] != "" {
//		dsn = MyEnv["DB_DSN"]
//	}
//	fmt.Println("dsn", dsn)
//	db, err := gorm.Open("mysql", dsn)
//	db.SingularTable(true)
//	db.SetLogger(EsErrlogger)
//	if err != nil {
//		panic(err)
//	}
//	db_max_idle, _ := strconv.Atoi(MyEnv["DB_MAX_IDLE"])
//	db_max_open, _ := strconv.Atoi(MyEnv["DB_MAX_OPEN"])
//	db.DB().SetMaxIdleConns(db_max_idle)
//	db.DB().SetMaxOpenConns(db_max_open)
//	if MyEnv["APP_DEBUG"] == "TRUE" {
//		db.LogMode(true)
//	}
//	return db
}
