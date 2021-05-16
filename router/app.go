package router

import (
	"ginShoppingMall/app/controller/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * APP 接口
 */

func AppRouter(router *gin.Engine) *gin.Engine {

	apis := router.Group("/api/")
	{
		// 用户/企业 接口
		user := apis.Group("user/")
		{
			// 用户/企业 注册
			user.POST("register", app.Register)
			// 用户/企业 登录
			user.POST("login", app.Login)
			// 用户/企业 刷新 token
			user.POST("authorize/refresh", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "login3333")
			})
		}

		// 支付 接口

		// 图库 接口

		// 其他 接口
	}

	return router
}
