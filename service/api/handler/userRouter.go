package handler

import (
	auth "boxin/service/auth/proto/auth"
	user "boxin/service/user/proto/user"
	verify "boxin/service/verification/proto/verification"
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
		v1.POST("/getinfo", getinfo)   //获取个人信息
		v1.POST("/info", editinfo)     //修改个人信息
		v1.POST("/password", setPw)    //修改登录密码
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
		Name     string `form:"name" json:"name" binding:"required"`
	}
	var p param

	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== register username======")
	log.Println(p.UserName)

	//验证码检验
	codeparam := verify.VerifyCodeEmailParam{
		Email:    p.Email,
		Username: p.UserName,
		Code:     p.Authcode}
	result1, err1 := verifyService.VerifyCodeEmail(context.Background(), &codeparam)
	log.Println(result1)
	log.Println(err1)
	if err1 != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "验证码验证过程有误，请稍后再试"})
		return
	}
	if result1.Status != 0 {
		c.JSON(200, gin.H{"status": 400, "msg": "验证码错误，请输入正确的验证码"})
		return
	}
	u := user.RegisterUserParam{
		UserName: p.UserName,
		Password: p.Password,
		School:   p.School,
		ID:       p.ID,
		Phone:    p.Phone,
		Email:    p.Email,
		Name:     p.Name}
	result, err := userService.RegisterStudent(context.Background(), &u)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "写入数据库失败", "data": err})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": result.Msg})
	return

}

func getinfo(c *gin.Context) {
	type userinfo struct {
		UserID   int32  `form:"userId" json:"userId"  binding:"required"`
		UserType int32  `form:"userType" json:"userType" binding:"required"`
		UserName string `form:"userName" json:"userName" binding:"required"`
		School   string `form:"school" json:"school" binding:"required"`
		ID       string `form:"ID" json:"ID"  binding:"required"`
		Phone    string `form:"phone" json:"phone" binding:"required"`
		Email    string `form:"email" json:"email"  binding:"required"`
		Name     string `form:"name" json:"name" binding:"required"`
	}
	type response struct {
		User userinfo `form:"user" json:"user" binding:"required"`
	}
	token, err := c.Cookie("token")
	log.Println(err)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	log.Println("====== userhandler——>getinfo token======")
	log.Println(token)
	//检验token
	ck := auth.CheckAuthParam{
		Token: token}
	info, jwterr := authService.CheckAuth(context.Background(), &ck)
	log.Println(info)
	log.Println(jwterr)
	if jwterr != nil {
		c.JSON(200, gin.H{"status": 404, "msg": "token失效，请重新登录"})
		return
	}

	userID := user.UserID{
		UserID: info.Data.UserID}

	result, err := userService.SearchUser(context.Background(), &userID)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取用户信息失败", "data": err})
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
			Name:     result.User.Name},
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取用户信息成功", "data": res})
	return
}

func editinfo(c *gin.Context) {
	type param struct {
		Phone    string `form:"phone" json:"phone" binding:"required"`
		Email    string `form:"email" json:"email"  binding:"required"`
		Password string `form:"password" json:"password"  binding:"required"`
	}

	token, err := c.Cookie("token")
	log.Println(err)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	log.Println("====== userhandler——>editinfo token======")
	log.Println(token)

	//检验token
	ck := auth.CheckAuthParam{
		Token: token}
	info, jwterr := authService.CheckAuth(context.Background(), &ck)
	log.Println(info)
	log.Println(jwterr)
	if jwterr != nil {
		c.JSON(200, gin.H{"status": 404, "msg": "token失效，请重新登录"})
		return
	}

	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== editinfo userId======")
	log.Println(info.Data.UserID)
	if info.Data.Password != p.Password {
		c.JSON(200, gin.H{"status": 505, "msg": "密码错误,请确认后重试"})
		return
	}
	uID := user.UserID{
		UserID: info.Data.UserID}
	data, err := userService.SearchUser(context.Background(), &uID)
	log.Println(data)
	log.Println(err)
	userinfo := data.User

	updateParam := user.UpdateUserParam{
		UserID:   userinfo.UserID,
		UserName: userinfo.UserName,
		UserType: userinfo.UserType,
		Password: userinfo.Password,
		School:   userinfo.School,
		ID:       userinfo.ID,
		Phone:    p.Phone,
		Email:    p.Email,
		Name:     userinfo.Name}

	result, err := userService.UpdateUser(context.Background(), &updateParam)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库更新失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "修改个人用户信息成功"})
}

func setPw(c *gin.Context) {
	type param struct {
		NewPW string `form:"newPw" json:"newPw"  binding:"required"`
		OldPW string `form:"oldPw" json:"oldPw"  binding:"required"`
	}
	token, err := c.Cookie("token")
	log.Println(err)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}

	//解析检验token
	log.Println("====== userhandler——>set password token======")
	log.Println(token)
	ck := auth.CheckAuthParam{
		Token: token}
	usrinfo, jwterr := authService.CheckAuth(context.Background(), &ck)
	log.Println(usrinfo)
	log.Println(jwterr)
	if jwterr != nil {
		c.JSON(200, gin.H{"status": 404, "msg": "token失效，请重新登录"})
		return
	}

	//获取参数
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	if p.OldPW != usrinfo.Data.Password {
		c.JSON(200, gin.H{"status": 505, "msg": "密码错误,请确认后重试"})
		return
	}
	if p.OldPW == p.NewPW {
		a := auth.LogoutParam{
			Token: token}
		resultOut, errOut := authService.Logout(context.Background(), &a)
		log.Println(resultOut)
		log.Println(errOut)
		c.JSON(200, gin.H{"status": 200, "msg": "新旧密码一致,当前token已失效，请重新登录"})
		return
	}
	uID := user.UserID{
		UserID: usrinfo.Data.UserID}
	data, err := userService.SearchUser(context.Background(), &uID)
	log.Println(data)
	log.Println(err)
	info := data.User

	updateParam := user.UpdateUserParam{
		UserID:   info.UserID,
		UserName: info.UserName,
		UserType: info.UserType,
		Password: p.NewPW,
		School:   info.School,
		ID:       info.ID,
		Phone:    info.Phone,
		Email:    info.Email,
		Name:     info.Name}

	result, err := userService.UpdateUser(context.Background(), &updateParam)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库更新失败"})
		return
	}
	a := auth.LogoutParam{
		Token: token}
	resultOut, errOut := authService.Logout(context.Background(), &a)
	log.Println(resultOut)
	log.Println(errOut)
	c.JSON(200, gin.H{"status": 200, "msg": "密码重置成功"})
}
