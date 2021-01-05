/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-17 10:20:03
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-05 22:42:42
 */
package handler

import (
	auth "boxin/service/auth/proto/auth"
	courseclass "boxin/service/courseclass/proto/courseclass"
	utils "boxin/utils"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var courseClassService courseclass.CourseClassService

//CourseRouter 处理与班级有关的api
func CourseRouter(g *gin.Engine, s courseclass.CourseClassService) {
	courseClassService = s
	v1 := g.Group("/course")
	{
		v1.GET("/mylist", getmylist)    //获取个人课程列表
		v1.GET("/student", getstudent)  //获取个人信息
		v1.PUT("/newcourse", newcourse) //新增课程
		v1.POST("/edit", editcourse)    //编辑课程
	}
}

func getmylist(c *gin.Context) {
	type param struct {
		UserID int32 `form:"userId" json:"userId"  binding:"required"`
	}
	token, err1 := c.Cookie("token")
	log.Println(err1)
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
		c.JSON(200, gin.H{"status": 404, "msg": "token失效，请重新登录", "data": jwterr})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== getmylist userId======")
	log.Println(p.UserID)
	ID := courseclass.UserID{
		UserID: p.UserID}
	var result *courseclass.SearchTakeByUserResponse
	var err error
	//教师
	if usrinfo.Data.UserType == 1 {

	} else {
		result, err = courseClassService.SearchTakeByUser(context.Background(), &ID)
	}

	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": result.Msg, "data": result.Courses})
}

func getstudent(c *gin.Context) {
	type param struct {
		CourseID int32 `form:"courseId" json:"courseId"  binding:"required"`
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== getCourseStudent CoursId======")
	log.Println(p.CourseID)
	ID := courseclass.CourseID{
		CourseID: p.CourseID,
	}
	result, err := courseClassService.SearchTakeByCourse(context.Background(), &ID)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败或未找到相应数据"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": result.Msg, "data": result.Users})
}

func newcourse(c *gin.Context) {
	type param struct {
		CourseName   string `form:"courseName" json:"courseName" binding:"required"`
		Introduction string `form:"introduction" json:"introduction" binding:"required"`
		Textbooks    string `form:"textbooks" json:"textbooks" binding:"required"`
		StartTime    string `form:"startTime" json:"startTime" binding:"required"`
		EndTime      string `form:"endTime" json:"endTime" binding:"required"`
	}
	type res struct {
		CourseID int32 `form:"courseId" json:"courseId" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== courseRouter——>newCourse token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有创建课程的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== newCourse CourseName======")
	log.Println(p.CourseName)

	startTime := utils.String2timeStamp(p.StartTime)
	endTime := utils.String2timeStamp(p.EndTime)
	log.Println(startTime)
	log.Println(endTime)
	log.Println(utils.TimeStamp2string(startTime)) //输出：2019-01-08 13:50:30
	log.Println(utils.TimeStamp2string(endTime))   //输出：2019-01-08 13:50:30
	newC := courseclass.NewCourseMessage{
		UserID:       usrinfo.Data.UserID,
		CourseName:   p.CourseName,
		Introduction: p.Introduction,
		TextBooks:    p.Textbooks,
		StartTime:    startTime,
		EndTime:      endTime,
		State:        1}
	result, err := courseClassService.NewCourse(context.Background(), &newC)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	resdata := res{CourseID: result.Courseclass.CourseID}
	c.JSON(200, gin.H{"status": 200, "msg": "课程开设成功", "data": resdata})
}

func editcourse(c *gin.Context) {
	type param struct {
		CourseID     int32  `form:"courseId" json:"courseId" binding:"required"`
		CourseName   string `form:"courseName" json:"courseName" binding:"required"`
		Introduction string `form:"introduction" json:"introduction" binding:"required"`
		Textbooks    string `form:"textbooks" json:"textbooks" binding:"required"`
		StartTime    string `form:"startTime" json:"startTime" binding:"required"`
		EndTime      string `form:"endTime" json:"endTime" binding:"required"`
		State        int32  `form:"state" json:"state" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== courseRouter——>editCourse token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有编辑课程的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== newCourse CourseName======")
	log.Println(p.CourseName)

	startTime := utils.String2timeStamp(p.StartTime)
	endTime := utils.String2timeStamp(p.EndTime)
	log.Println(startTime)
	log.Println(endTime)

	log.Println(utils.TimeStamp2string(startTime)) //输出：2019-01-08 13:50:30
	log.Println(utils.TimeStamp2string(endTime))   //输出：2019-01-08 13:50:30
	newC := courseclass.CourseClass{
		CourseID:     p.CourseID,
		CourseName:   p.CourseName,
		Introduction: p.Introduction,
		TextBooks:    p.Textbooks,
		StartTime:    startTime,
		EndTime:      endTime,
		State:        p.State}
	result, err := courseClassService.UpdateCourseClass(context.Background(), &newC)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "课程信息修改成功"})
}
