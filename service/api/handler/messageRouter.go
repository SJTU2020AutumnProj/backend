/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-15 21:26:29
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 14:30:10
 */
package handler

import (
	auth "boxin/service/auth/proto/auth"
	message "boxin/service/message/proto/message"
	"boxin/utils"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var messageService message.MessageService

//checkRouter 注册answer有关接口
func MessageRouter(g *gin.Engine, s message.MessageService) {
	messageService = s
	v1 := g.Group("/message")
	{
		v1.GET("/userId", getPersonalMessage) //获取用户message
		// v1.POST("/info", editinfo)    //修改个人信息
	}
}

func getPersonalMessage(c *gin.Context) {
	type msgInfo struct {
		MsgID    int32  `form:"msgId" json:"msgId"  binding:"required"`
		MsgTime  string `form:"msgTime" json:"msgTime"  binding:"required"`
		MsgType  int32  `form:"msgType" json:"msgType"  binding:"required"`
		UserID   int32  `form:"userId" json:"userId"  binding:"required"`
		CourseID int32  `form:"courseId" json:"courseId"  binding:"required"`
		Title    string `form:"title" json:"title"  binding:"required"`
		Content  string `form:"content" json:"content"  binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== MessageRouter——>getPersonalMessage token======")
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

	log.Println("====== getPersonalMessage UserID======")
	log.Println(usrinfo.Data.UserID)

	commitTime := time.Now().Unix()
	log.Println(commitTime)
	log.Println(utils.TimeStamp2string2(commitTime))
	a := message.GetMessageByUserIDParam{
		UserID: usrinfo.Data.UserID,
	}

	result, err := messageService.GetMessageByUserID(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	resdata := make([]msgInfo, len(result.Data))
	for i, v := range result.Data {
		resdata[i] = msgInfo{
			MsgID:    v.MessageID,
			MsgTime:  utils.TimeStamp2string2(v.MessageTime),
			MsgType:  v.MessageType,
			UserID:   v.UserID,
			CourseID: v.CourseID,
			Title:    v.Title,
			Content:  v.Content,
		}
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取个人信息", "data": resdata})

}
