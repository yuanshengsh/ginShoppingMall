package app

import (
	"errors"
	"ginShoppingMall/app/dto"
	"ginShoppingMall/app/model"
	"ginShoppingMall/app/service"
	boot "ginShoppingMall/bootstrap"
	"ginShoppingMall/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserRegister
// @Router /api/user/register [post]
// @Summary 用户注册
// @Description 用户注册
// @Tags API USER 用户
// @Param mobile path int true "手机号"
// @Param name path string true "用户名"
// @Param password path string true "密码"
// @Param captcha path string true "手机验证码"
func UserRegister(c *gin.Context) {
	params := &dto.UserRegisterInput{}
	if err := params.BindValidParam(c); err != nil {
		util.ResponseError(c, 2000, err)
		return
	}

	user := &model.User{}

	// 检查手机号格式
	mobile := strconv.Itoa(params.Mobile)
	if false == util.ValidatorMobile(mobile) {
		util.ResponseError(c, 2000, errors.New("不是正确的手机号"))
		return
	}

	// 检查手机号是否已经注册过
	model.FindUserByMobile(mobile, user)
	if user.ID != 0 {
		util.ResponseError(c, 2000, errors.New("该手机号已被注册"))
		return
	}

	user.Mobile = mobile
	user.Name = params.Name
	user.Password = util.GenSaltPassword(boot.Config.Server.Salt, params.Password)

	if err := model.SaveUser(user); err != nil {
		util.ResponseError(c, 2000, err)
		return
	}

	util.ResponseSuccess(c, "注册成功")
}

// UserLogin
// @Router /api/user/login [post]
// @Summary 用户登录
// @Description 用户登录
// @Tags API USER 用户
// @Param name path string true "用户名"
// @Param password path string true "密码"
func UserLogin(c *gin.Context) {
	params := &dto.UserLoginInput{}
	if err := params.BindValidParam(c); err != nil {
		util.ResponseError(c, 2000, err)
		return
	}

	user := &model.User{}

	password := util.GenSaltPassword(boot.Config.Server.Salt, params.Password)
	// 检查用户
	model.FindUserLogin(params.Name, password, user)
	if user.ID == 0 {
		util.ResponseError(c, 2000, errors.New("账号错误，登录失败"))
		return
	}

	ut := &model.UserToken{}
	token, _ := service.GenToken(user.ID, user.Mobile)
	ut.Token = token
	ut.UserId = user.ID
	if err := model.SaveUserToken(ut); err != nil {
		util.ResponseError(c, 2000, err)
		return
	}

	output := dto.UserLoginOutput{
		AccessToken: "Bearer " + token,
		ExpiresAt:   ut.ExpiresAt,
	}
	util.ResponseSuccess(c, output)
}

// GetUser
// @Router /api/user [get]
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags API USER 用户
func GetUser(c *gin.Context) {
	user_name, _ := c.Get("user_name")
	util.ResponseSuccess(c, user_name)
}

// PutUser
// @Router /api/user [put]
// @Summary 获取用户信息
// @Description 获取用户信息
// @Tags API USER 用户
func PutUser(c *gin.Context) {
	util.ResponseSuccess(c, "PutUser")
}
