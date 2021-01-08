/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-12-15 21:26:29
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-07 16:51:06
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
}
