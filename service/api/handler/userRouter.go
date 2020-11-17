package handler

import (
	user "boxin/service/user/proto/user"

	"github.com/gin-gonic/gin"
)

var userService user.UserService

func UserRouter(g *gin.Engine, s user.UserService) {
	userService = s
	v1 := g.Group("/user")
	{
		v1.POST("/register", register) //注册
		v1.GET("/info", getinfo)       //获取个人信息
		v1.POST("/info", editinfo)     //修改个人信息
	}
}

func register(c *gin.Context) {

	type param struct {
		userName string `form:"userName" binding:"required"`
		password string `form:"password" binding:"required"`
		school   string `form:"school" binding:"required"`
		ID       int64  `form:"ID" binding:"required"`
		phone    string `form:"phone" binding:"required"`
		email    string `form:"email" binding:"required"`
		authcode string `form:"authcode" binding:"rquried"`
	}
	type response struct {
		status int16  `JSON:"status"`
		msg    string `JSON:"msg"`
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		res := &response{500, "缺少必须参数，请稍后重试"}
		c.JSON(200, res)
		return
	}
	if p.authcode != "123456" {
		res := &response{401, "验证码有误，请重新确认后再试"}
		c.JSON(200, res)
		return
	}
	u := user.User{
		UserType: 1,
		UserName: p.userName,
		Password: p.password,
		School:   p.school,
		Phone:    p.phone,
		Email:    p.email}
	result, err := userService.AddUser(c, &u)
	if err != nil {
		res := &response{401, "写入数据库有误"}
		c.JSON(200, res)
		return
	}
	// res := &response{200, "写入成功"}
	c.JSON(200, result)
	return

}

func getinfo(c *gin.Context) {

}

func editinfo(c *gin.Context) {

}
