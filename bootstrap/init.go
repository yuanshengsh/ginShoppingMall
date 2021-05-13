package bootstrap

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

/**
 * 校验初始化配置，校验 MySQL Redis Log
 * 配置全局 Config
 */

var Config = &Cfg{}

func init() {
	pwd, _ := os.Getwd()
	InitCommonConfig(pwd + "/../../config/config.yaml")
}

type Cfg struct {
	App   Server `yaml:"App"`
	Admin Server `yaml:"Admin"`
	//MySQL  MySQL
	//Redis  Redis
}

type Server struct {
	Port                    string `yaml:"Port"`
	AllowCrossDomain        bool   `yaml:"AllowCrossDomain"`
	LogPath                 string `yaml:"LogPath"`
	LogType                 string `yaml:"LogType"`
	LogInfoName             string `yaml:"LogInfoName"`
	LogErrorName            string `yaml:"LogErrorName"`
	JwtTokenSignKey         string `yaml:"JwtTokenSignKey"`
	JwtTokenOnlineUsers     int    `yaml:"JwtTokenOnlineUsers"`
	JwtTokenCreatedExpireAt int    `yaml:"JwtTokenCreatedExpireAt"`
	JwtTokenRefreshExpireAt int    `yaml:"JwtTokenRefreshExpireAt"`
}

type Log struct {
	Path      string `yaml:"Path"`
	Type      string `yaml:"Type"`
	InfoName  string `yaml:"InfoName"`
	ErrorName string `yaml:"ErrorName"`
}

//type MySQL struct {
//}
//
//type Redis struct {
//}

func InitCommonConfig(path string) {
	//根据路径读文件内容
	content, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}

	//yaml解析，并赋值给Config
	if err = yaml.Unmarshal(content, &Config); err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
}
