package util

import (
	"fmt"
	boot "ginShoppingMall/bootstrap"
	"github.com/gomodule/redigo/redis"
	"os"
	"strconv"
	"time"
)

func StartRedis() *redis.Pool {
	redisMaxIdle, _ := strconv.Atoi(boot.Config.Redis.MaxIdle)
	redisMaxActive, _ := strconv.Atoi(boot.Config.Redis.MaxActive)
	redisIdleTimeout, _ := strconv.Atoi(boot.Config.Redis.IdleTimeout)
	redisHost, redisPort, redisDb := boot.Config.Redis.Host, boot.Config.Redis.Port, boot.Config.Redis.IndexDb
	redisOption := redis.DialPassword(boot.Config.Redis.Auth)

	RedisClient := &redis.Pool{
		MaxIdle:     redisMaxIdle,
		MaxActive:   redisMaxActive,
		IdleTimeout: time.Duration(redisIdleTimeout),
		Dial: func() (redis.Conn, error) {
			client, err := redis.Dial("tcp", redisHost+":"+redisPort, redisOption)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(-1)
			}
			client.Do("SELECT", redisDb)
			return client, nil
		},
	}
	return RedisClient
}
