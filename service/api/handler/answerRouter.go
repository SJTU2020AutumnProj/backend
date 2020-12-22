package handler

// import (
// 	answer "boxin/service/answer/proto/answer"
// 	repo "boxin/service/answer/repository"
// 	user "boxin/service/user/proto/user"
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/gin-gonic/gin"
// )

// var answerService answer.AnswerService

// //AnswerRouter 注册answer有关接口
// func AnswerRouter(g *gin.Engine, s user.UserService) {
// 	userService = s
// 	v1 := g.Group("/answer")
// 	{
// 		v1.POST("/new", CreateAnswer) //学生提交答案
// 		// v1.GET("/info", getinfo)      //获取个人信息
// 		// v1.POST("/info", editinfo)    //修改个人信息
// 	}
// }

// // createAnswer
// func CreateAnswer(c *gin.Context) {
// 	type param struct {
// 		HwID       int32  `form:"hwId" json:"hwId" binding:"required"`
// 		StuID      int32  `form:"stuId" json:"stuId"  binding:"required"`
// 		Status     int32  `form:"status" json:"status"  binding:"required"`
// 		CommitTime string `form:"commitTime" json:"commitTime" binding:"required"`
// 	}
// 	var p param

// 	if err := c.ShouldBindJSON(&p); err != nil {
// 		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
// 		return
// 	}
// 	log.Println("====== create Answer======")
// 	log.Println(p.HwID)

// 	timeTemplate1 := "2019-01-08 13:50:30" //常规类型
// 	ctime, _ := time.ParseInLocation(timeTemplate1, p.CommitTime, time.Local)
// 	commitTime := ctime.Unix()
// 	log.Println(commitTime)
// 	a := answerService.Answer{
// 		HomeworkID: p.HwID,
// 		StudentID:  p.StuID,
// 		Status:     p.Status,
// 		CommitTime: ctime,
// 	}
// 	result, err := answerService.CreateAnswer(context.Background(), &a)
// 	log.Println(result)
// 	log.Println(err)
// 	if err != nil {
// 		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
// 		return
// 	}

// 	c.JSON(200, gin.H{"status": 200, "msg": result.Msg})

// }

// func getinfo(c *gin.Context) {
// 	type param struct {
// 		UserID int32 `form:"userId" json:"userId"  binding:"required"`
// 		// UserID int32 `form:"userId" binding:"required"`
// 	}

// 	type userinfo struct {
// 		UserID   int32  `form:"userId" json:"userId"  binding:"required"`
// 		UserType int32  `form:"userType" json:"userType" binding:"required"`
// 		UserName string `form:"userName" json:"userName" binding:"required"`
// 		School   string `form:"school" json:"school" binding:"required"`
// 		ID       string `form:"ID" json:"ID"  binding:"required"`
// 		Phone    string `form:"phone" json:"phone" binding:"required"`
// 		Email    string `form:"email" json:"email"  binding:"required"`
// 	}
// 	type response struct {
// 		User userinfo `form:"user" json:"user" binding:"required"`
// 	}
// 	var p param
// 	if err := c.ShouldBind(&p); err != nil {
// 		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
// 		return
// 	}
// 	log.Println("====== getinfo userId======")
// 	log.Println(p.UserID)
// 	ID := user.UserID{
// 		UserID: p.UserID}
// 	result, err := userService.SearchUser(context.Background(), &ID)
// 	log.Println(result)
// 	log.Println(err)
// 	if err != nil {
// 		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
// 		return
// 	}
// 	res := response{
// 		User: userinfo{
// 			UserID:   result.User.UserID,
// 			UserName: result.User.UserName,
// 			UserType: result.User.UserType,
// 			School:   result.User.School,
// 			ID:       result.User.Id,
// 			Phone:    result.User.Phone,
// 			Email:    result.User.Email,
// 		},
// 	}
// 	c.JSON(200, gin.H{"status": 200, "msg": "获取信息成功", "data": res})
// 	return
// }

// func editinfo(c *gin.Context) {
// 	type userparam struct {
// 		UserID   int32  `form:"userId" json:"userId"  binding:"required"`
// 		UserName string `form:"userName" json:"userName" binding:"required"`
// 		School   string `form:"school" json:"school" binding:"required"`
// 		ID       string `form:"ID" json:"ID"  binding:"required"`
// 		Phone    string `form:"phone" json:"phone" binding:"required"`
// 		Email    string `form:"email" json:"email"  binding:"required"`
// 	}
// 	type param struct {
// 		User userparam `form:"user" binding:"required"`
// 	}
// 	var p param
// 	if err := c.ShouldBind(&p); err != nil {
// 		log.Println(err)
// 		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
// 		return
// 	}
// 	log.Println("====== editinfo userId======")
// 	log.Println(p.User.UserID)
// 	ID := user.UserID{
// 		UserID: p.User.UserID}
// 	data, err := userService.SearchUser(context.Background(), &ID)

// 	user := data.User

// 	var flag bool
// 	flag = false
// 	if user.UserName != p.User.UserName {
// 		user.UserName = p.User.UserName
// 		flag = true
// 	}
// 	if user.School != p.User.School {
// 		user.School = p.User.School
// 		flag = true
// 	}
// 	if user.Id != p.User.ID {
// 		user.Id = p.User.ID
// 		flag = true
// 	}
// 	if user.Phone != p.User.Phone {
// 		user.Phone = p.User.Phone
// 		flag = true
// 	}

// 	if user.Email != p.User.Email {
// 		user.Email = p.User.Email
// 		flag = true
// 	}
// 	if flag == false {
// 		c.JSON(200, gin.H{"status": 200, "msg": "未发生改动，修改信息与数据库信息一致"})
// 		return
// 	}
// 	result, err := userService.UpdateUser(context.Background(), user)
// 	log.Println(result)
// 	log.Println(err)
// 	if err != nil {
// 		c.JSON(200, gin.H{"status": 401, "msg": "数据库更新失败"})
// 		return
// 	}
// 	c.JSON(200, gin.H{"status": 200, "msg": result.Msg})
// }
