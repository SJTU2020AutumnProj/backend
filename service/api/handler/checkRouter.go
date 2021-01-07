/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-15 21:26:29
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 11:36:30
 */
package handler

import (
	check "boxin/service/check/proto/check"

	"github.com/gin-gonic/gin"
)

var checkService check.CheckService

//checkRouter 注册answer有关接口
func CheckRouter(g *gin.Engine, s check.CheckService) {
	checkService = s
	// v1 := g.Group("/check")
	// {
	// 	v1.PUT("/correct", teacherCheck) //j教师批改作业
	// 	v1.GET("/info", getinfo)      //获取个人信息
	// 	v1.POST("/info", editinfo)    //修改个人信息
	// }
}
