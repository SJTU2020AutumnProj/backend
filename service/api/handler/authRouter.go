/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-18 08:38:54
 * @LastEditors: Seven
 * @LastEditTime: 2020-11-18 23:36:12
 */
package handler

import (
	auth "boxin/service/auth/proto/auth"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var authService auth.AuthService

//AuthRouter 注册user有关接口
func AuthRouter(g *gin.Engine, s auth.AuthService) {
	authService = s
	v1 := g.Group("/auth")
	{
		v1.POST("/login", login)  //登录
		v1.GET("/logout", logout) //退出登录

	}
}

func login(c *gin.Context) {
	type param struct {
		UserName string `form:"userName" json:"userName" binding:"required"`
		Password string `form:"password" json:"password"  binding:"required"`
	}
	type resdata struct {
		User auth.UserData `form:"user" json:"user"`
	}

	var p param

	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== login username======")
	log.Println(p.UserName)
	var a auth.LoginParam
	a.UserName = p.UserName
	a.Password = p.Password

	result, err := authService.Login(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": err})
		return
	}
	user := resdata{
		User: *result.Data}
	c.JSON(200, gin.H{"status": 200, "msg": "登录成功", "data": user})
	return

}

func logout(c *gin.Context) {

	c.JSON(200, gin.H{"status": 200, "msg": "退出登录成功"})
	return
}
