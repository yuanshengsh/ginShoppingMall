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

// Register
// @Router /api/user/register [post]
// @Summary 用户注册
// @Description 用户注册
// @Tags API USER 用户
// @Param mobile path int false "手机号"
// @Param name path string false "用户名"
// @Param password path string false "密码"
// @Param captcha path string false "手机验证码"
func Register(c *gin.Context) {
	params := &dto.RegisterInput{}
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

// Login
// @Router /api/user/login [post]
// @Summary 用户登录
// @Description 用户登录
// @Tags API USER 用户
// @Param name path int false "用户名"
// @Param password path int false "密码"
func Login(c *gin.Context) {
	params := &dto.LoginInput{}
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

	output := dto.LoginOutput{
		AccessToken: "Bearer " + token,
		ExpiresAt:   ut.ExpiresAt,
	}
	util.ResponseSuccess(c, output)
}
