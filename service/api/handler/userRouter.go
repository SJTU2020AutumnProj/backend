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
		UserName string `form:"userName" json:"userName" binding:"required"`
		Password string `form:"password" json:"password"  binding:"required"`
		School   string `form:"school" json:"school"  binding:"required"`
		ID       string `form:"ID" json:"ID" binding:"required"`
		Phone    string `form:"phone" json:"phone" binding:"required"`
		Email    string `form:"email" json:"email"  binding:"required"`
		Authcode string `form:"authcode" json:"authcode" binding:"required"`
	}
	var p param

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== register username======")
	log.Println(p.UserName)
	if p.Authcode != "123456" {
		c.JSON(200, gin.H{"status": 401, "msg": "验证码有误，请重新确认后再试"})
		return
	}
	u := user.RegisterUserParam{
		UserName: p.UserName,
		Password: p.Password,
		School:   p.School,
		ID:       p.ID,
		Phone:    p.Phone,
		Email:    p.Email}

	result, err := userService.RegisterStudent(context.Background(), &u)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "写入数据库失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": result.Msg})
	return

}

func getinfo(c *gin.Context) {
	type param struct {
		UserID int32 `form:"userId" json:"userId"  binding:"required"`
		// UserID int32 `form:"userId" binding:"required"`
	}

	type userinfo struct {
		UserID   int32  `form:"userId" json:"userId"  binding:"required"`
		UserType int32  `form:"userType" json:"userType" binding:"required"`
		UserName string `form:"userName" json:"userName" binding:"required"`
		School   string `form:"school" json:"school" binding:"required"`
		ID       string `form:"ID" json:"ID"  binding:"required"`
		Phone    string `form:"phone" json:"phone" binding:"required"`
		Email    string `form:"email" json:"email"  binding:"required"`
	}
	type response struct {
		User userinfo `form:"user" json:"user" binding:"required"`
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
	result, err := userService.SearchUser(context.Background(), &ID)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	res := response{
		User: userinfo{
			UserID:   result.User.UserID,
			UserName: result.User.UserName,
			UserType: result.User.UserType,
			School:   result.User.School,
			ID:       result.User.ID,
			Phone:    result.User.Phone,
			Email:    result.User.Email,
		},
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取信息成功", "data": res})
	return
}

func editinfo(c *gin.Context) {
	type userparam struct {
		UserID   int32  `form:"userId" json:"userId"  binding:"required"`
		UserName string `form:"userName" json:"userName" binding:"required"`
		School   string `form:"school" json:"school" binding:"required"`
		ID       string `form:"ID" json:"ID"  binding:"required"`
		Phone    string `form:"phone" json:"phone" binding:"required"`
		Email    string `form:"email" json:"email"  binding:"required"`
	}
	type param struct {
		User userparam `form:"user" binding:"required"`
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== editinfo userId======")
	log.Println(p.User.UserID)
	ID := user.UserID{
		UserID: p.User.UserID}
	data, err := userService.SearchUser(context.Background(), &ID)
	log.Println(data)
	log.Println(err)
	user := data.User

	var flag bool
	flag = false
	if user.UserName != p.User.UserName {
		user.UserName = p.User.UserName
		flag = true
	}
	if user.School != p.User.School {
		user.School = p.User.School
		flag = true
	}
	if user.ID != p.User.ID {
		user.ID = p.User.ID
		flag = true
	}
	if user.Phone != p.User.Phone {
		user.Phone = p.User.Phone
		flag = true
	}

	if user.Email != p.User.Email {
		user.Email = p.User.Email
		flag = true
	}
	if flag == false {
		c.JSON(200, gin.H{"status": 200, "msg": "未发生改动，修改信息与数据库信息一致"})
		return
	}

	// result, err := userService.UpdateUser(context.Background(), &user)
	// log.Println(result)
	// log.Println(err)
	// if err != nil {
	// 	c.JSON(200, gin.H{"status": 401, "msg": "数据库更新失败"})
	// 	return
	// }
	// c.JSON(200, gin.H{"status": 200, "msg": result.Msg})
}
