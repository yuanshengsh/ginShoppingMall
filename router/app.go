package router

import (
	"ginShoppingMall/app/controller/app"
	"ginShoppingMall/middleware"
	"github.com/gin-gonic/gin"
)

/**
 * APP 接口
 */

func AppRouter(router *gin.Engine) *gin.Engine {

	// 不需要授权的接口
	router.POST("/api/user/register", app.UserRegister) // 用户注册
	router.POST("/api/user/login", app.UserLogin)       // 用户登录
	router.GET("/api/captcha/pic", app.CaptchaPic)      // 获取图片验证码
	router.POST("/api/captcha/sms", app.CaptchaSMS)     // 获取短信验证码

	apis := router.Group("/api/", middleware.JwtApp())
	{
		// 用户接口
		user := apis.Group("user/")
		{
			//// 获取用户信息
			user.GET("info", app.UserInfo)
		//	// 更新用户信息
		//	user.GET("info", func(ctx *gin.Context) {
		//		name, _ := ctx.Get("user_name")
		//		ctx.String(http.StatusOK, "login3333")
		//	})
		}

		// 支付 接口

		// 图库 接口

		// 其他 接口
	}

	return router
}
