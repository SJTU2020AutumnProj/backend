/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-15 21:26:29
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-05 15:02:32
 */
package handler

// import (
// 	// "boxin/service/answer/proto/answer"
// 	answer "boxin/service/answer/proto/answer"
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
// 	a := answer.CreateAnswerParam{
// 		HomeworkID: p.HwID,
// 		StudentID:  p.StuID,
// 		Status:     p.Status,
// 		CommitTime: commitTime,
// 	}
// 	result, err := answerService.CreateAnswer(context.Background(), &a)
// 	log.Println(result)
// 	log.Println(err)
// 	if err != nil {
// 		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
// 		return
// 	}

// 	c.JSON(200, gin.H{"status": 200, "msg": result.Msg, "answerId": result.AnswerID})

// }
