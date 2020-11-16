package handler

import (
	user "boxin/service/user/proto/user"

	"github.com/gin-gonic/gin"
)

var userService user.UserService

func UserRouter(g *gin.Engine, s user.UserService) {
	userService = s
	v1 := g.Group("/user")
	{
		v1.GET("/")
	}
}
