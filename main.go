package main

import (
	"github.com/gin-gonic/gin"
)

// @title ginShoppingMall API
// @version 1.0
// @description 接口描述
// @termsOfService http://swagger.io/terms/

// @contact.name Cheney
// @contact.url http://adaday.cn
// @contact.email YuanShengSH@outlook.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host :8080
// @BasePath /api
func main() {
	r := gin.Default()
	//url := ginSwagger.URL("docs/swagger.json")
	r.Run(":8080")
}
