/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-15 21:26:29
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 09:31:00
 */
package handler

import (
	answer "boxin/service/answer/proto/answer"
	auth "boxin/service/auth/proto/auth"
	"boxin/utils"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var answerService answer.AnswerService

//AnswerRouter 注册answer有关接口
func AnswerRouter(g *gin.Engine, s answer.AnswerService) {
	answerService = s
	v1 := g.Group("/answer")
	{
		v1.PUT("/new", stuCreateAnswer) //学生提交答案
		// v1.GET("/info", getinfo)      //获取个人信息
		// v1.POST("/info", editinfo)    //修改个人信息
	}
}

// createAnswer
func stuCreateAnswer(c *gin.Context) {
	type param struct {
		HwID       int32  `form:"hwId" json:"hwId" binding:"required"`
		Content    string `form:"content" json:"content"  binding:"required"`
		CommitTime string `form:"commitTime" json:"commitTime" binding:"required"`
		Note       string `form:"note" json:"note" `
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== AnswerRouter——>stuCreateAnswer token======")
	log.Println(token)
	ck := auth.CheckAuthParam{
		Token: token}
	usrinfo, jwterr := authService.CheckAuth(context.Background(), &ck)
	log.Println(usrinfo)
	log.Println(jwterr)
	if jwterr != nil {
		c.JSON(200, gin.H{"status": 404, "msg": "token失效，请重新登录", "data": jwterr})
		return
	}
	//不是学生？
	if usrinfo.Data.UserType != 1 {
		c.JSON(200, gin.H{"status": 500, "msg": "您应该调用教师上传答案接口！"})
		return
	}
	var p param

	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== stuCreateAnswer======")
	log.Println(p.HwID)

	commitTime := utils.String2timeStamp2(p.CommitTime)
	log.Println(commitTime)
	log.Println(utils.TimeStamp2string2(commitTime))
	a := answer.PostAnswerParam{
		HomeworkID: p.HwID,
		UserID:     usrinfo.Data.UserID,
		Content:    p.Content,
		CommitTime: commitTime,
		Note:       p.Note,
	}

	result, err := answerService.PostAnswerByStudent(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "msg": result.Msg, "answerId": result.AnswerID})

}
