package handler

import (
	user "boxin/service/user/proto/user"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var userService user.UserService

//UserRouter 注册user有关接口
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
		ID       string `form:"ID"  binding:"required"`
		Phone    string `form:"phone" binding:"required"`
		Email    string `form:"email"  binding:"required"`
		Authcode string `form:"authcode"  binding:"required"`
	}
	var p param

	if err := c.ShouldBind(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== register username======")
	log.Println(p.UserName)
	if p.Authcode != "123456" {
		c.JSON(200, gin.H{"status": 401, "msg": "验证码有误，请重新确认后再试"})
		return
	}
	u := user.UserInfo{
		UserName: p.UserName,
		Password: p.Password,
		School:   p.School,
		Id:       p.ID,
		Phone:    p.Phone,
		Email:    p.Email}
	result, err := userService.RegisterStudent(context.Background(), &u)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "写入数据库有误"})
		return
	}
	c.JSON(200, result)
	return

}

func getinfo(c *gin.Context) {
	type param struct {
		UserID int32 `form:"userId"  binding:"required"`
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== getinfo userId======")
	log.Println(p.UserID)
	ID := user.UserID{
		UserID: p.UserID}
	result, err := userService.SearchUser(c, &ID)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, result)

}

func editinfo(c *gin.Context) {
	type user struct {
		UserID   int32  `form:"userId"  binding:"required"`
		UserName string `form:"userName"  binding:"required"`
		UserType int32  `form:"userType" binding:"required"`
		Password string `form:"password"  binding:"required"`
		School   string `form:"school"  binding:"required"`
		ID       int64  `form:"ID"  binding:"required"`
		Phone    string `form:"phone" binding:"required"`
		Email    string `form:"email"  binding:"required"`
	}
	type param struct {
		User user `form:"user" binding:"required"`
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== editinfo userId======")
	log.Println(p.User.UserID)

	// result, err := userService.SearchUser(c, &ID)
	// if err != nil {
	// 	c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
	// 	return
	// }
	// c.JSON(200, result)
}
