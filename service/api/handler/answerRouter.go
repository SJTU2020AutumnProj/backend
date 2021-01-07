/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-15 21:26:29
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 16:51:33
 */
package handler

import (
	answer "boxin/service/answer/proto/answer"
	auth "boxin/service/auth/proto/auth"
	check "boxin/service/check/proto/check"
	"boxin/utils"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var answerService answer.AnswerService

//AnswerRouter 注册answer有关接口
func AnswerRouter(g *gin.Engine, s answer.AnswerService) {
	answerService = s
	v1 := g.Group("/answer")
	{
		v1.PUT("/new", stuCreateAnswer)         //学生提交答案
		v1.POST("/get", getAnswer)              //获取学生作业答案
		v1.POST("/checkdetail", getCheckDetail) //获取批改作业信息
		// v1.POST("/info", editinfo)    //修改个人信息
	}
}

// createAnswer
func stuCreateAnswer(c *gin.Context) {
	type param struct {
		HwID    int32  `form:"hwId" json:"hwId" binding:"required"`
		Content string `form:"content" json:"content"  binding:"required"`
		Note    string `form:"note" json:"note" `
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
	if usrinfo.Data.UserType != 0 {
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

	commitTime := time.Now().Unix()
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

// 获取学生作业答案
func getAnswer(c *gin.Context) {
	type param struct {
		AnswerID int32 `form:"answerId" json:"answerId"  binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== AnswerRouter——>getAnswer token======")
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
	//不是教师？
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
	log.Println("====== GetAnswer======")
	log.Println(p.AnswerID)

	commitTime := time.Now().Unix()
	log.Println(commitTime)
	log.Println(utils.TimeStamp2string2(commitTime))
	a := answer.AnswerID{
		AnswerID: p.AnswerID,
	}

	result, err := answerService.SearchAnswer(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "msg": "获取答案内容成功", "answerId": result.Answer})

}

func getCheckDetail(c *gin.Context) {
	type param struct {
		CheckID int32 `form:"checkId" json:"checkId"  binding:"required"`
	}
	type CheckInfo struct {
		CheckID     int32  `form:"checkId" json:"checkId"  binding:"required"`
		CheckTime   string `form:"checkTime" json:"checkTime"  binding:"required"`
		Description string `form:"description" json:"description"  binding:"required"`
		Comment     string `form:"comment" json:"comment"  binding:"required"`
		Score       int32  `form:"score" json:"score"  binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== AnswerRouter——>getCheckDetail token======")
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
	var p param

	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== getCheckDetail======")
	log.Println(p.CheckID)

	commitTime := time.Now().Unix()
	log.Println(commitTime)
	log.Println(utils.TimeStamp2string2(commitTime))
	a := check.CheckID{
		CheckID: p.CheckID,
	}

	result, err := checkService.SearchCheckByID(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	resdata := CheckInfo{
		CheckID:     result.Check.CheckID,
		CheckTime:   utils.TimeStamp2string2(result.Check.CheckTime),
		Description: result.Check.Description,
		Comment:     result.Check.Comment,
		Score:       result.Check.Score,
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取批改内容成功", "data": resdata})

}
