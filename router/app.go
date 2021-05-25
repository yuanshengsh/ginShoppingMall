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

	apis := router.Group("/api", middleware.JwtApp())
	{
		// 用户信息
		apis.GET("/user", app.GetUser) // 获取用户信息
		apis.PUT("/user", app.PutUser) // 更新用户信息

		// 用户收货地址
		apis.GET("/user/address", app.PutUser)    // 获取用户收货地址
		apis.GET("/user/address/1", app.PutUser)  // 获取收货地址详情
		apis.POST("/user/address/1", app.PutUser) // 添加收货地址
		apis.PUT("/user/address/1", app.PutUser)  // 更新收货地址

		// 商品相关
		apis.GET("/product/:ID", app.GetProduct)  // 获取商品信息
		//apis.GET("/product/spec", app.GetProduct) // 获取商品规格

		// 支付 接口

		// 图库 接口

		// 其他 接口
	}

	return router
}
