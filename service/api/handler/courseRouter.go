/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-17 10:20:03
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-06 13:32:13
 */
package handler

import (
	auth "boxin/service/auth/proto/auth"
	courseclass "boxin/service/courseclass/proto/courseclass"
	homework "boxin/service/homework/proto/homework"
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
		v1.GET("/mylist", getmylist)       //获取个人课程列表
		v1.GET("/student", getstudent)     //获取个人信息
		v1.PUT("/newcourse", newcourse)    //新增课程
		v1.POST("/edit", editcourse)       //编辑课程
		v1.GET("/detail", courseDetail)    //编辑课程
		v1.POST("/delete", courseDelete)   //编辑课程
		v1.GET("/hwlist", getHWlist)       //获取作业列表
		v1.GET("/nostu", getNotinStudent)  //获取作业列表
		v1.POST("/students", addStudents)  //课程添加学生
		v1.POST("/delstu", deleteStudents) //课程添加学生
	}
}

func getmylist(c *gin.Context) {
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
	log.Println("====== getmylist userId======")
	log.Println(usrinfo.Data.UserID)
	ID := courseclass.UserID{
		UserID: usrinfo.Data.UserID}
	result, err := courseClassService.SearchTakeByUser(context.Background(), &ID)

	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "查询成功", "data": result.Courses})
}

func getstudent(c *gin.Context) {
	type param struct {
		CourseID int32 `form:"courseId" json:"courseId"  binding:"required"`
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
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
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
	c.JSON(200, gin.H{"status": 200, "msg": "获取学生列表成功", "data": result.Users})
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

func courseDetail(c *gin.Context) {
	type param struct {
		CourseID int32 `form:"courseId" json:"courseId"  binding:"required"`
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
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== getCourseDetail CoursId======")
	log.Println(p.CourseID)
	ID := courseclass.CourseID{
		CourseID: p.CourseID,
	}
	result, err := courseClassService.SearchCourseClass(context.Background(), &ID)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败或未找到相应数据"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取学生列表成功", "data": result.Courseclass})

}

func courseDelete(c *gin.Context) {
	type param struct {
		CourseID int32 `form:"courseId" json:"courseId"  binding:"required"`
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有删除课程的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== newCourse CourseName======")
	log.Println(p.CourseID)

	a := courseclass.CourseID{
		CourseID: p.CourseID}
	result, err := courseClassService.DeleteCourseClass(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "课程信息删除成功"})
}

func getHWlist(c *gin.Context) {
	type param struct {
		CourseID int32 `form:"courseId" json:"courseId"  binding:"required"`
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
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== getHwlist CoursId======")
	log.Println(p.CourseID)
	ID := homework.CourseID{
		CourseID: p.CourseID,
	}
	result, err := homeworkService.SearchHomeworkByCourseID(context.Background(), &ID)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败或未找到相应数据"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取作业列表成功", "data": result.Homeworks})
	return
}

func getNotinStudent(c *gin.Context) {
	type param struct {
		CourseID int32 `form:"courseId" json:"courseId"  binding:"required"`
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
	//不是教师？
	if usrinfo.Data.UserType != 1 {
		c.JSON(200, gin.H{"status": 500, "msg": "您没有创建课程的权限！"})
		return
	}
	var p param
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== getNotinStudent CoursId======")
	log.Println(p.CourseID)
	// ID := homework.CourseID{
	// 	CourseID: p.CourseID,
	// }
	// result, err := homeworkService.SearchHomeworkByCourseID(context.Background(), &ID)
	// log.Println(result)
	// log.Println(err)
	// if err != nil {
	// 	c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败或未找到相应数据"})
	// 	return
	// }
	c.JSON(200, gin.H{"status": 200, "msg": "获取学生列表成功"})
	return
}

func addStudents(c *gin.Context) {
	// stus1 := c.PostForm("students")
	// log.Println(stus1)
	type studentID struct {
		UserID int32 `form:"userId" json:"userId" binding:"required"`
	}
	type param struct {
		Students []studentID `form:"students" json:"students" binding:"required"`
		CourseID int32       `form:"courseId" json:"courseId"  binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== courseRouter——>addStudents token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有添加学生的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}

	log.Println("====== addStudents CourseName======")
	log.Println(p.CourseID)

	students := make([]int32, len(p.Students))
	for i, v := range p.Students {
		students[i] = v.UserID
	}

	stus := courseclass.Take{
		UserID:   students[0],
		CourseID: p.CourseID,
		Role:     0}
	result, err := courseClassService.AddTake(context.Background(), &stus)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "重复添加或学生不存在"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "msg": "学生添加成功"})
}

func deleteStudents(c *gin.Context) {
	// stus1 := c.PostForm("students")
	// log.Println(stus1)
	type studentID struct {
		UserID int32 `form:"userId" json:"userId" binding:"required"`
	}
	type param struct {
		Students []studentID `form:"students" json:"students" binding:"required"`
		CourseID int32       `form:"courseId" json:"courseId"  binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== courseRouter——>addStudents token======")
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

	log.Println("====== addStudents CourseName======")
	log.Println(p.CourseID)

	students := make([]int32, len(p.Students))
	for i, v := range p.Students {
		students[i] = v.UserID
	}

	stus := courseclass.UserCourse{
		UserID:   students[0],
		CourseID: p.CourseID}
	result, err := courseClassService.DeleteTake(context.Background(), &stus)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "重复添加或学生不存在"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "学生删除成功"})
}
