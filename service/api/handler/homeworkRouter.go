/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2021-01-06 10:11:40
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-06 23:56:21
 */
package handler

import (
	auth "boxin/service/auth/proto/auth"
	homework "boxin/service/homework/proto/homework"
	utils "boxin/utils"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var homeworkService homework.HomeworkService

//HomeworkRouter 注册作业有关接口
func HomeworkRouter(g *gin.Engine, s homework.HomeworkService) {
	homeworkService = s
	v1 := g.Group("/homework")
	{
		v1.PUT("/create", createHW)     //创建作业
		v1.POST("/update", modifyHw)    //修改作业
		v1.GET("/detail", stuGetdetail) //退出登录
		// v1.GET("/checkAuth", checkAuth) //检测权限
	}
}

func createHW(c *gin.Context) {
	type param struct {
		Title       string `form:"title" json:"title" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Note        string `form:"note" json:"note" binding:"required"`
		Content     string `form:"content" json:"content" binding:"required"`
		CourseID    int32  `form:"courseId" json:"courseId" binding:"required"`
		State       int32  `form:"state" json:"state" binding:"required"`
		//0表示暂存，未发布，1表示发布
		Score     int32  `form:"score" json:"score" binding:"required"`
		StartTime string `form:"startTime" json:"startTime" binding:"required"`
		EndTime   string `form:"endTime" json:"endTime" binding:"required"`
	}
	token, err := c.Cookie("token")
	log.Println(err)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	log.Println("====== logout token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有创建作业的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== newHomework Title======")
	log.Println(p.Title)
	startTime := utils.String2timeStamp(p.StartTime)
	endTime := utils.String2timeStamp(p.EndTime)
	log.Println(startTime)
	log.Println(endTime)
	log.Println(utils.TimeStamp2string(startTime)) //输出：2019-01-08 13:50:30
	log.Println(utils.TimeStamp2string(endTime))   //输出：2019-01-08 13:50:30
	newHW := homework.AssignHomeworkParam{
		UserID:      usrinfo.Data.UserID,
		CourseID:    p.CourseID,
		Title:       p.Title,
		State:       p.State,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: p.Description,
		Content:     p.Content,
		Note:        p.Note,
		Score:       p.Score,
	}
	result, err := homeworkService.AssignHomework(context.Background(), &newHW)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "新建作业成功", "data": result.HomeworkID})
	return
}

func modifyHw(c *gin.Context) {
	type param struct {
		HwID        int32  `form:"hwId" json:"hwId" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Note        string `form:"note" json:"note" binding:"required"`
		Content     string `form:"content" json:"content" binding:"required"`
		CourseID    int32  `form:"courseId" json:"courseId" binding:"required"`
		State       int32  `form:"state" json:"state" binding:"required"`
		//0表示暂存，未发布，1表示发布
		Score     int32  `form:"score" json:"score" binding:"required"`
		StartTime string `form:"startTime" json:"startTime" binding:"required"`
		EndTime   string `form:"endTime" json:"endTime" binding:"required"`
	}
	token, err := c.Cookie("token")
	log.Println(err)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	log.Println("====== logout token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有修订作业的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== newHomework Title======")
	log.Println(p.Title)
	startTime := utils.String2timeStamp(p.StartTime)
	endTime := utils.String2timeStamp(p.EndTime)
	log.Println(startTime)
	log.Println(endTime)
	log.Println(utils.TimeStamp2string(startTime)) //输出：2019-01-08 13:50:30
	log.Println(utils.TimeStamp2string(endTime))   //输出：2019-01-08 13:50:30
	newHW := homework.HomeworkInfo{
		UserID:      usrinfo.Data.UserID,
		CourseID:    p.CourseID,
		Title:       p.Title,
		State:       p.State,
		StartTime:   startTime,
		EndTime:     endTime,
		Description: p.Description,
		Content:     p.Content,
		Note:        p.Note,
		HomeworkID:  p.HwID,
		Score:       p.Score,
	}
	result, err := homeworkService.UpdateHomework(context.Background(), &newHW)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, gin.H{"status": 200, "msg": "修改作业成功"})
	return
}

func stuGetdetail(c *gin.Context) {
	type param struct {
		HwID int32 `form:"hwId" json:"hwId" binding:"required"`
	}
	type resdata struct {
		HwID        int32  `form:"hwId" json:"hwId" binding:"required"`
		Title       string `form:"title" json:"title" binding:"required"`
		Description string `form:"description" json:"description" binding:"required"`
		Note        string `form:"note" json:"note" binding:"required"`
		Content     string `form:"content" json:"content" binding:"required"`
		CourseID    int32  `form:"courseId" json:"courseId" binding:"required"`
		State       int32  `form:"state" json:"state" binding:"required"`
		//0表示暂存，未发布，1表示发布
		Score            int32  `form:"score" json:"score" binding:"required"`
		StartTime        string `form:"startTime" json:"startTime" binding:"required"`
		EndTime          string `form:"endTime" json:"endTime" binding:"required"`
		AnswerID         int32  `form:"answerId " json:"answerId" binding:"required"`
		CheckID          int32  `form:"checkId" json:"checkId" binding:"required"`
		StandardAnswerID int32  `form:"standardAnswerId" json:"standardAnswerId" binding:"required"`
		TeacherID        int32  `form:"teacherId" json:"teacherId" binding:"required"`
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
	var p param
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== stuGetdetail hwId======")
	log.Println(p.HwID)
	userHw := homework.GetUserHomeworkParam{
		UserID:     usrinfo.Data.UserID,
		HomeworkID: p.HwID,
	}
	res1, err := homeworkService.GetUserHomework(context.Background(), &userHw)

	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	usrhw := res1.UserHomework
	log.Println(usrhw)

	hwid := homework.HomeworkID{
		HomeworkID: p.HwID,
	}
	res2, err := homeworkService.SearchHomework(context.Background(), &hwid)

	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	hwinfo := res2.Homework
	log.Println(hwinfo)

	// res:=resdata{
	// 	HwID :hwinfo.HomeworkID,
	// 	Title :hwinfo.Title,
	// 	Description :hwinfo.HomeworkID,
	// 	Note    :hwinfo.HomeworkID,
	// 	Content    :hwinfo.HomeworkID,
	// 	CourseID    int32  `form:"courseId" json:"courseId" binding:"required"`
	// 	State       int32  `form:"state" json:"state" binding:"required"`
	// 	//0表示暂存，未发布，1表示发布
	// 	Score     int32  `form:"score" json:"score" binding:"required"`
	// 	StartTime string `form:"startTime" json:"startTime" binding:"required"`
	// 	EndTime   string `form:"endTime" json:"endTime" binding:"required"`
	// 	AnswerID    int32  `form:"answerId " json:"answerId" binding:"required"`
	// 	CheckID    int32  `form:"checkId" json:"checkId" binding:"required"`
	// 	StandardAnswerID    int32  `form:"standardAnswerId" json:"standardAnswerId" binding:"required"`
	// 	TeacherID
	// }
	// c.JSON(200, gin.H{"status": 200, "msg": "课程开设成功", "data": result.})

}
