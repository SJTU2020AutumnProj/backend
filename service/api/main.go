/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-16 21:32:52
 * @LastEditors: Seven
 * @LastEditTime: 2020-12-15 19:17:22
 */
package main

import (
	"boxin/service/api/handler"
	auth "boxin/service/auth/proto/auth"
	courseclass "boxin/service/courseclass/proto/courseclass"
	user "boxin/service/user/proto/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

const (
	// AppName     = "go.micro.service.api"
	ServiceName = "go.micro.api.api"
	EtcdAddr    = "localhost:2379"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		fmt.Println(method)
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法，因为有的模板是要请求两次的
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}

func main() {
	etcdRegister := etcd.NewRegistry(
		registry.Addrs(EtcdAddr),
	)
	// app := micro.NewService(
	// 	micro.Name(AppName),
	// 	micro.Registry(etcdRegister))
	userService := user.NewUserService("go.micro.service.user", client.DefaultClient)
	authService := auth.NewAuthService("go.micro.service.auth", client.DefaultClient)
	courseClassService := courseclass.NewCourseClassService("go.micro.service.courseclass", client.DefaultClient)
	webHandler := gin.Default()
	webHandler.Use(Cors())
	service := web.NewService(
		web.Name(ServiceName),
		web.Address(":8080"),
		web.Handler(webHandler),
		web.Registry(etcdRegister),
	)
	handler.UserRouter(webHandler, userService)
	handler.AuthRouter(webHandler, authService)
	handler.CourseRouter(webHandler, courseClassService)
	// handler.CourseRouter(webHandler,courService)
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
