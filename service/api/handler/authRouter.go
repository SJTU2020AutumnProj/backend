/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-18 08:38:54
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 16:51:20
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
		v1.POST("/login", login)         //登录
		v1.POST("/logout", logout)       //退出登录
		v1.POST("/checkAuth", checkAuth) //检测权限
	}
}

func login(c *gin.Context) {
	type param struct {
		UserName string `form:"userName" json:"userName" binding:"required"`
		Password string `form:"password" json:"password"  binding:"required"`
	}
	type resdata struct {
		User auth.UserData `form:"user" json:"user"`
		// Token string        `form:"token" json:"token"`
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
		c.JSON(200, gin.H{"status": 401, "msg": "登录失败", "data": err})
		return
	}
	c.SetCookie("token", result.Token, 3600, "/", "/", false, true)
	user := resdata{
		User: *result.Data}
	// Token: result.Token}
	c.JSON(200, gin.H{"status": 200, "msg": "登录成功", "data": user})
	return

}

func logout(c *gin.Context) {
	token, err := c.Cookie("token")
	log.Println(err)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	log.Println("====== logout token======")
	log.Println(token)
	a := auth.LogoutParam{
		Token: token}
	result, err := authService.Logout(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": err})
		return
	}
	//退出登录，重置cookie
	c.SetCookie("token", "", 1, "/", "/", false, true)
	c.JSON(200, gin.H{"status": 200, "msg": "退出登录成功"})
	return
}

func checkAuth(c *gin.Context) {
	type Resdata struct {
		UserID   int32  `form:"userId" json:"userId" binding:"required"`
		UserType int32  `form:"userType" json:"userType"  binding:"required"`
		UserName string `form:"userName" json:"userName" binding:"required"`
	}

	// 获取cookie参数
	token, err := c.Cookie("token")
	log.Println(err)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
	}
	log.Println("====== checkAuth token======")
	log.Println(token)
	ck := auth.CheckAuthParam{
		Token: token}
	result, err := authService.CheckAuth(context.Background(), &ck)
	log.Println(result)
	log.Println(err)
	if nil != err {
		c.JSON(200, gin.H{"status": 401, "msg": err})
		return
	}
	res := Resdata{
		UserID:   result.Data.UserID,
		UserType: result.Data.UserType,
		UserName: result.Data.UserName}

	c.JSON(200, gin.H{"status": 200, "msg": "权限验证成功", "data": res})
	return
}
