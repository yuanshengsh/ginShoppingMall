package dto

import (
	"ginShoppingMall/util"
	"github.com/gin-gonic/gin"
)

type UserRegisterInput struct {
	Mobile   int    `json:"mobile" comment:"手机号" example:"13999999999" validate:"required"`
	Name     string `json:"name" comment:"用户名" example:"username" validate:"required"`
	Password string `json:"password" comment:"密码" example:"abc@123!" validate:"required"`
}

func (param *UserRegisterInput) BindValidParam(c *gin.Context) error {
	return util.DefaultGetValidParams(c, param)
}

type UserLoginInput struct {
	Name     string `json:"name" comment:"用户名" example:"username" validate:"required"`
	Password string `json:"password" comment:"密码" example:"abc@123!" validate:"required"`
}

func (param *UserLoginInput) BindValidParam(c *gin.Context) error {
	return util.DefaultGetValidParams(c, param)
}

type UserLoginOutput struct {
	AccessToken string `json:"access_token" form:"access_token"` //access_token
	ExpiresAt   string `json:"expires_at" form:"expires_at"`     //expires_in
}
