/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-17 10:20:03
 * @LastEditors: Seven
 * @LastEditTime: 2020-11-17 21:05:41
 */
package handler

import (
	user "boxin/service/user/proto/user"
	"log"

	"github.com/gin-gonic/gin"
)

func CourseRouter(g *gin.Engine) {
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
	log.Println("====== getinfo userId======")
	log.Println(p.UserID)
	ID := user.UserID{
		UserID: p.UserID}
	result, err := userService.SearchUser(c, &ID)
	if err != nil {
		c.JSON(200, gin.H{"status": 401, "msg": "数据库读取失败"})
		return
	}
	c.JSON(200, result)
}

func getstudent(c *gin.Context) {

}

func newcourse(c *gin.Context) {

}
