package handler

import (
	"github.com/gin-gonic/gin"
)

func courseRouter(g *gin.Engine) {
	v1 := g.Group("/course")
	{
		v1.GET("/mylist", getmylist)    //获取个人课程列表
		v1.GET("/student", getstudent)  //获取个人信息
		v1.PUT("/newcourse", newcourse) //新增课程
	}
}

func getmylist(c *gin.Context) {

}

func getstudent(c *gin.Context) {

}

func newcourse(c *gin.Context) {

}
