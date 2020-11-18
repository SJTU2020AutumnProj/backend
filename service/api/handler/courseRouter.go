/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-17 10:20:03
 * @LastEditors: Seven
 * @LastEditTime: 2020-11-18 20:53:49
 */
package handler

import (
	courseclass "boxin/service/courseclass/proto/courseclass"
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
	}
}

func getmylist(c *gin.Context) {
	type param struct {
		UserID int32 `form:"userId" json:"userId"  binding:"required"`
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
	result, err := courseClassService.SearchTakeByUser(context.Background(), &ID)
	log.Println(result)
	log.Println(err)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, result)
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
	c.JSON(200, result)
}

func newcourse(c *gin.Context) {
	type course struct {
		CourseName   string `form:"courseName" json:"courseName" binding:"required"`
		Introduction string `form:"introduction" json:"introduction" binding:"required"`
		Textbooks    string `form:"textbooks" json:"textbooks" binding:"required"`
		StartTime    int64  `form:"startTime" json:"startTime" binding:"required"`
		EndTime      int64  `form:"endTime" json:"endTime" binding:"required"`
	}
	type param struct {
		Course course `form:"course" json:"course" binding:"required"`
		UserId int32  `form:"userId" json:"userId" binding:"required"`
	}
	var p param
	if err := c.ShouldBind(&p); err != nil {
		c.JSON(200, gin.H{"status": 500, "msg": "缺少必须参数，请稍后重试"})
		return
	}
	log.Println("====== newCourse CourseName======")
	log.Println(p.Course.CourseName)
	// teacher := courseclass.UserID{
	// 	UserID: p.UserId,
	// }
	// newC := courseclass.CourseClass{
	// 	CourseName:   p.Course.CourseName,
	// 	Introduction: p.Course.Introduction,
	// 	TextBooks:    p.Course.Textbooks,
	// 	StartTime:    p.Course.StartTime,
	// 	EndTime:      p.Course.EndTime,
	// }
	// result, err := courseClassService.AddCourseClass(context.Background(), &newC)

	// log.Println(result)
	// log.Println(err)
	// if err != nil {
	// 	c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
	// 	return
	// }

	// c.JSON(200, gin.H{"status": 200})
}
