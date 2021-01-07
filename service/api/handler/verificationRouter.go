/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2021-01-05 15:49:05
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-08 00:00:42
 */
package handler

import (
	verify "boxin/service/verification/proto/verification"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var verifyService verify.VerificationService

//VerifyRouter 注册验证码有关接口
func VerifyRouter(g *gin.Engine, s verify.VerificationService) {
	verifyService = s
	v1 := g.Group("/verification")
	{
		v1.POST("/send", sendCodeEmail) //发送验证码
	}
}

func sendCodeEmail(c *gin.Context) {
	type param struct {
		UserName string `form:"userName" json:"userName" binding:"required"`
		Email    string `form:"email" json:"email"  binding:"required"`
	}
	var p param
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("======verify_send username======")
	log.Println(p.UserName)
	a := verify.SendCodeEmailParam{
		Email:    p.Email,
		Username: p.UserName}
	result, err := verifyService.SendCodeEmail(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "邮箱发送异常"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "验证码已发送"})
	return
}
