package app

import (
	"ginShoppingMall/util"
	"github.com/gin-gonic/gin"
)

// Product
// @Router /api/product [get]
// @Summary 商品信息
// @Description 商品信息
// @Tags API Product 商品
func Product(c *gin.Context) {
	util.ResponseSuccess(c, "Product")
}
