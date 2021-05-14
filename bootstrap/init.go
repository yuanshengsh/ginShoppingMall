package bootstrap

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

/**
 * 校验初始化配置，校验 MySQL Redis Log
 * 配置全局 Config
 */

var (
	Config = &Cfg{}
	DB     *gorm.DB
	RP     *redis.Pool
)

type Cfg struct {
	Server Server `yaml:"Server"`
	Log    Log    `yaml:"Log"`
	Token  Token  `yaml:"Token"`
	MySQL  MySQL  `yaml:"MySQL"`
	Redis  Redis  `yaml:"Redis"`
}

type Server struct {
	Port             string `yaml:"Port"`
	AllowCrossDomain bool   `yaml:"AllowCrossDomain"`
}

type Log struct {
	Path      string `yaml:"Path"`
	Type      string `yaml:"Type"`
	InfoName  string `yaml:"InfoName"`
	ErrorName string `yaml:"ErrorName"`
}

type Token struct {
	JwtTokenSignKey         string `yaml:"JwtTokenSignKey"`
	JwtTokenOnlineUsers     int    `yaml:"JwtTokenOnlineUsers"`
	JwtTokenCreatedExpireAt int    `yaml:"JwtTokenCreatedExpireAt"`
	JwtTokenRefreshExpireAt int    `yaml:"JwtTokenRefreshExpireAt"`
}

type MySQL struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Pass     string `yaml:"Pass"`
	DataBase string `yaml:"DataBase"`
	Prefix   string `yaml:"Prefix"`
}

type Redis struct {
	Host        string `yaml:"Host"`
	Port        string `yaml:"Port"`
	Auth        string `yaml:"Auth"`
	IndexDb     string `yaml:"IndexDb"`
	MaxIdle     string `yaml:"MaxIdle"`
	MaxActive   string `yaml:"MaxActive"`
	IdleTimeout string `yaml:"IdleTimeout"`
}

func InitCommonConfig(env string) {
	// 获取config目录地址
	pwd, _ := os.Getwd()
	path := pwd + "/../../config/" + env + ".yaml"

	// 根据路径读文件内容
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	// yaml解析，并赋值给Config
	if err = yaml.Unmarshal(content, &Config); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
