package dto

import (
	"ginShoppingMall/util"
	"github.com/gin-gonic/gin"
)

type CaptchaPicInput struct {
	Mobile int  `bson:"mobile" validate:"required"`
	//Mobile int `json:"mobile" comment:"手机号" validate:"required"`
}

func (param *CaptchaPicInput) BindValidParam(c *gin.Context) error {
	return util.DefaultGetValidParams(c, param)
}
