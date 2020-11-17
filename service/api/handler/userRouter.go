package handler

import (
	user "boxin/service/user/proto/user"
	"log"

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
		UserName string `form:"userName"  binding:"required"`
		Password string `form:"password"  binding:"required"`
		School   string `form:"school"  binding:"required"`
		ID       int64  `form:"ID"  binding:"required"`
		Phone    string `form:"phone" binding:"required"`
		Email    string `form:"email"  binding:"required"`
		Authcode string `form:"authcode"  binding:"required"`
	}
	type response struct {
		status int16  `JSON:"status"`
		msg    string `JSON:"msg"`
	}
	var p param
	// if err := c.ShouldBind(&p); err != nil {
	// 	c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
	// 	return
	// }
	// log.Println("====== Bind by query String ======")
	// log.Println(p.UserName)
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== Bind By JSON ======")
	log.Println(p.UserName)
	if p.Authcode != "123456" {
		c.JSON(200, gin.H{"status": 401, "msg": "验证码有误，请重新确认后再试"})
		return
	}
	u := user.User{
		UserType: 1,
		UserName: p.UserName,
		Password: p.Password,
		School:   p.School,
		Phone:    p.Phone,
		Email:    p.Email}
	result, err := userService.AddUser(c, &u)
	if err != nil {

		c.JSON(200, gin.H{"status": 401, "msg": "写入数据库有误"})
		return
	}
	c.JSON(200, result)
	return

}

func getinfo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "err"})
}

func editinfo(c *gin.Context) {

}
