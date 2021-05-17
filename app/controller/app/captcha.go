package app

import (
	"ginShoppingMall/app/dto"
	boot "ginShoppingMall/bootstrap"
	"ginShoppingMall/util"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"image/color"
	"image/png"
)

const CAPTCHA_PIC_TTL = 600 // 验证码生命周期 10 分钟

// CaptchaPic
// @Router /api/captcha/pic [get]
// @Summary 获取图片验证码
// @Description 获取图片验证码
// @Tags API USER 用户
// @Param mobile query int true "手机号"
func CaptchaPic(c *gin.Context) {
	params := &dto.CaptchaPicInput{}
	if err := params.BindValidParam(c); err != nil {
		util.ResponseError(c, 2000, err)
		return
	}

	cap := captcha.New()
	//通过句柄调用 字体文件
	if err := cap.SetFont(boot.Config.Server.Font); err != nil {
		panic(err.Error())
	}

	//设置图片大小
	cap.SetSize(91, 50)
	//设置感染强度
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{R: 255, G: 255, B: 255, A: 255})
	cap.SetBkgColor(
		color.RGBA{R: 255, A: 255},
		color.RGBA{B: 255, A: 255},
		color.RGBA{G: 153, A: 255},
	)

	c.Header("Content-Type", "image/png")

	img, str := cap.Create(4, captcha.NUM)
	util.RedisCall("SETEX", params.Mobile, CAPTCHA_PIC_TTL, str)
	png.Encode(c.Writer, img)
}

// CaptchaSMS
// @Router /api/captcha/sms [post]
// @Summary 获取短信验证码
// @Description 获取短信验证码
// @Tags API USER 用户
// @Param mobile path int true "手机号"
// @Param captcha path int true "图形验证码"
func CaptchaSMS(c *gin.Context) {
	//user_name := c.Get("user_name")
	util.ResponseSuccess(c, "user_name")
}
