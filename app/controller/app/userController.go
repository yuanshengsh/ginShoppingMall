package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Register
// @Router /api/user/register [post]
// @Summary 用户/企业 注册
// @Description 用户/企业 注册
// @Tags API USER 用户
// @Param mobile path int false "手机号"
// @Param name path string false "用户名"
// @Param password path string false "密码"
func Register(c *gin.Context) {
	c.String(200, "123")
}

// Login
// @Router /api/user/login [post]
// @Summary 用户/企业 登录
// @Description 用户/企业 登录
// @Tags API USER 用户
// @Param name path int false "用户名"
// @Param password path int false "密码"
func Login(c *gin.Context) {
	//user := service.GetUser()
	//fmt.Println(user)
	//
	//json := make(map[string]interface{}) //注意该结构接受的内容
	var param struct {
		Name     string `json:"name"`
		Password string `json:"password" binding:"required"`
	}
	c.BindJSON(&param)
	fmt.Println(param.Name)
	c.String(200, "123")

	//name := c.PostForm("name")
	//password := c.Query("password")
	//
	//fmt.Println(name)
	//fmt.Println(password)

	// 检查用户是否存在

	// 校验密码是否正确

	// 生成 token

	////判断用户是否存在
	////存在输出状态1
	////不存在创建用户，保存密码与用户名
	//Bool := Func.IsExist(name)
	//if Bool {
	//	//注册状态
	//	State["state"] = 1
	//	State["text"] = "此用户已存在！"
	//} else {
	//	//用户不存在即添加用户
	//	AddStruct(name, passwd)
	//	State["state"] = 1
	//	State["text"] = "注册成功！"
	//}

	//把状态码和注册状态返回给客户端
	//c.String(http.StatusOK, "~")
}
