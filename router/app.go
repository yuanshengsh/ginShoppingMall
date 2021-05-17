package router

import (
	"fmt"
	"ginShoppingMall/app/controller/app"
	"ginShoppingMall/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * APP 接口
 */

func AppRouter(router *gin.Engine) *gin.Engine {

	// 不需要授权的接口
	router.POST("/api/user/register", app.Register) // 用户注册
	router.POST("/api/user/login", app.Login)       // 用户登录
	router.GET("/api/captcha/pic", app.CaptchaPic)  // 图片验证码
	//router.POST("/api/captcha/sms", app.Login)      // 发送短信验证码

	apis := router.Group("/api/", middleware.JwtApp())
	{
		// 用户接口
		user := apis.Group("user/")
		{
			// 获取用户信息
			user.GET("info", func(ctx *gin.Context) {
				name, _ := ctx.Get("user_name")
				fmt.Println(name)
				ctx.String(http.StatusOK, "login3333")
			})
			// 更新用户信息
			user.POST("update", func(ctx *gin.Context) {
				name, _ := ctx.Get("user_name")
				fmt.Println(name)
				ctx.String(http.StatusOK, "login3333")
			})
			//// 用户/企业 刷新 token
			//user.POST("authorize/refresh", func(ctx *gin.Context) {
			//	ctx.String(http.StatusOK, "login3333")
			//})
		}

		// 支付 接口

		// 图库 接口

		// 其他 接口
	}

	return router
}
