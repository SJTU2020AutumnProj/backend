/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2021-01-06 10:11:40
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 16:50:28
 */
package handler

import (
	auth "boxin/service/auth/proto/auth"
	check "boxin/service/check/proto/check"
	homework "boxin/service/homework/proto/homework"
	utils "boxin/utils"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var homeworkService homework.HomeworkService

//HomeworkRouter 注册作业有关接口
func HomeworkRouter(g *gin.Engine, s homework.HomeworkService) {
	homeworkService = s
	v1 := g.Group("/homework")
	{
		v1.PUT("/create", createHW)                        //创建作业
		v1.POST("/update", modifyHw)                       //修改作业
		v1.POST("/detail", stuGetdetail)                   //退出登录
		v1.POST("/cmtlist", GetcmtList)                    //检测权限
		v1.PUT("/postAnswer", teacherPostAnswer)           //教师上传标准答案
		v1.POST("/publishAnswer", teacherPublishAnswer)    //教师发布标准答案
		v1.POST("/publishCheck", teacherPublishCheck)      //教师发布批改情况
		v1.PUT("/correct", teacherCheck)                   //教师批改作业
		v1.POST("/teacherGetHomework", teacherGetHomework) //教师批改作业

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
		UserID           int32  `form:"userId" json:"userId" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== HomeworkRouter——>stuGetdetail token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "请调用教师专用接口！"})
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
	userHwparam := homework.GetUserHomeworkParam{
		UserID:     usrinfo.Data.UserID,
		HomeworkID: p.HwID,
	}
	res1, err := homeworkService.GetUserHomework(context.Background(), &userHwparam)

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

	responsedata := resdata{
		HwID:             hwinfo.HomeworkID,
		Title:            hwinfo.Title,
		Description:      hwinfo.Description,
		Note:             hwinfo.Note,
		Content:          hwinfo.Content,
		CourseID:         hwinfo.CourseID,
		State:            usrhw.State,
		UserID:           usrhw.UserID,
		Score:            hwinfo.Score,
		StartTime:        utils.TimeStamp2string(hwinfo.StartTime),
		EndTime:          utils.TimeStamp2string(hwinfo.EndTime),
		AnswerID:         usrhw.AnswerID,
		CheckID:          usrhw.CheckID,
		StandardAnswerID: hwinfo.AnswerID,
		TeacherID:        hwinfo.UserID,
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取作业详情成功（学生）", "data": responsedata})

}

func GetcmtList(c *gin.Context) {
	type param struct {
		HwID int32 `form:"hwId" json:"hwId" binding:"required"`
	}
	type resdata struct {
		UserID   int32  `form:"userId" json:"userId"  binding:"required"`
		School   string `form:"school" json:"school" binding:"required"`
		ID       string `form:"ID" json:"ID"  binding:"required"`
		State    int32  `form:"state" json:"state"  binding:"required"`
		AnswerID int32  `form:"answerId" json:"answerId"  binding:"required"`
		CheckID  int32  `form:"checkId" json:"checkId"  binding:"required"`
		Name     string `form:"name" json:"name" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== HomeworkRouter——>GetcmtList token======")
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
	if err := c.ShouldBindJSON(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== GetcmtList hwId======")
	log.Println(p.HwID)
	hid := homework.HomeworkID{
		HomeworkID: p.HwID,
	}
	result, err := homeworkService.GetUserByHomeworkID(context.Background(), &hid)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	responsedata := make([]resdata, len(result.UserInfo))
	for i, v := range result.UserInfo {
		responsedata[i] = resdata{
			UserID:   v.UserID,
			School:   v.School,
			ID:       v.ID,
			State:    v.State,
			AnswerID: v.AnswerID,
			CheckID:  v.CheckID,
			Name:     v.Name,
		}
	}

	c.JSON(200, gin.H{"status": 200, "msg": "获取作业提交情况成功", "data": responsedata})

}

func teacherGetHomework(c *gin.Context) {
	type param struct {
		HwID int32 `form:"hwId" json:"hwId" binding:"required"`
	}
	type resdata struct {
		HwID             int32  `form:"hwId" json:"hwId"  binding:"required"`
		CourseID         int32  `form:"courseId" json:"courseId" binding:"required"`
		TeacherID        int32  `form:"teachereId" json:"teacherId" binding:"required"`
		StartTime        string `form:"startTime" json:"startTime" binding:"required"`
		EndTime          string `form:"endTime" json:"endTime" binding:"required"`
		Title            string `form:"title" json:"title" binding:"required"`
		Score            int32  `form:"score" json:"score" binding:"required"`
		StandardAnswerID int32  `form:"standardAnswerId" json:"standardAnswerId"  binding:"required"`
		Content          string `form:"content" json:"content" binding:"required"`
		Description      string `form:"description" json:"description" binding:"required"`
		Note             string `form:"note" json:"note" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== HomeworkRouter——>teacherGetdetail token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "请调用教师专用接口！"})
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

	responsedata := resdata{
		HwID:             hwinfo.HomeworkID,
		Title:            hwinfo.Title,
		Description:      hwinfo.Description,
		Note:             hwinfo.Note,
		Content:          hwinfo.Content,
		CourseID:         hwinfo.CourseID,
		Score:            hwinfo.Score,
		StartTime:        utils.TimeStamp2string(hwinfo.StartTime),
		EndTime:          utils.TimeStamp2string(hwinfo.EndTime),
		StandardAnswerID: hwinfo.AnswerID,
		TeacherID:        hwinfo.UserID,
	}
	c.JSON(200, gin.H{"status": 200, "msg": "获取作业详情成功（教师）", "data": responsedata})

}

func teacherPostAnswer(c *gin.Context) {
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
	log.Println("====== HomeworkRouter——>teacherPostAnswer token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有上传标准答案的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== teacherPostAnswer hwId======")
	log.Println(p.HwID)
	commitTime := time.Now().Unix()
	log.Println(commitTime)
	log.Println(utils.TimeStamp2string2(commitTime))
	a := homework.PostParam{
		UserID:     usrinfo.Data.UserID,
		HomeworkID: p.HwID,
		CommitTime: commitTime,
		Content:    p.Content,
		Note:       p.Note,
	}
	result, err := homeworkService.PostHomeworkAnswer(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "msg": "上传标准答案成功", "data": result.AnswerID})

}

func teacherPublishAnswer(c *gin.Context) {
	type param struct {
		HwID int32 `form:"hwId" json:"hwId" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== HomeworkRouter——>teacherPublicAnswer token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有上传标准答案的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== teacherPublicAnswer hwId======")
	log.Println(p.HwID)
	pubTime := time.Now().Unix()
	log.Println(pubTime)
	log.Println(utils.TimeStamp2string2(pubTime))
	a := homework.ReleaseParam{
		HomeworkID: p.HwID,
		TeacherID:  usrinfo.Data.UserID,
		PubTime:    pubTime,
	}
	result, err := homeworkService.ReleaseHomeworkAnswer(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "msg": "发布标准答案成功"})

}

func teacherCheck(c *gin.Context) {
	type param struct {
		AnswerID    int32  `form:"answerId" json:"answerId" binding:"required"`
		UserID      int32  `form:"userId" json:"userId" binding:"required"`
		HwID        int32  `form:"hwId" json:"hwId" binding:"required"`
		Score       int32  `form:"score" json:"score" binding:"required"`
		Description string `form:"description" json:"descripttion" binding:"required"`
		Comment     string `form:"comment" json:"comment" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== HomeworkRouter——>teacherPublicAnswer token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有上传标准答案的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== teacherCheck hwId======")
	log.Println(p.HwID)
	checkTime := time.Now().Unix()
	log.Println(checkTime)
	log.Println(utils.TimeStamp2string2(checkTime))
	a := check.CreateCheckParam{
		HomeworkID:  p.HwID,
		TeacherID:   usrinfo.Data.UserID,
		CheckTime:   checkTime,
		StudentID:   p.UserID,
		Description: p.Description,
		Comment:     p.Comment,
		Score:       p.Score,
		AnswerID:    p.AnswerID,
	}
	result, err := checkService.CreateCheck(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "msg": "批改作业成功", "data": result.CheckID})
}

func teacherPublishCheck(c *gin.Context) {
	type param struct {
		HwID int32 `form:"hwId" json:"hwId" binding:"required"`
	}
	//获取token
	token, err1 := c.Cookie("token")
	log.Println(err1)
	if token == "" {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少token，请检查是否已登录"})
		return
	}
	//解析检验token
	log.Println("====== HomeworkRouter——>teacherPublicAnswer token======")
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
		c.JSON(200, gin.H{"status": 500, "msg": "您没有上传标准答案的权限！"})
		return
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== teacherPublicAnswer hwId======")
	log.Println(p.HwID)
	pubTime := time.Now().Unix()
	log.Println(pubTime)
	log.Println(utils.TimeStamp2string2(pubTime))
	a := homework.ReleaseCheckParam{
		HomeworkID:  p.HwID,
		TeacherID:   usrinfo.Data.UserID,
		ReleaseTime: pubTime,
	}
	result, err := homeworkService.ReleaseCheck(context.Background(), &a)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}

	c.JSON(200, gin.H{"status": 200, "msg": "发布作业批改情况成功"})

}
