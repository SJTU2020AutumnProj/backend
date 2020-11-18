/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-16 21:32:52
 * @LastEditors: Seven
 * @LastEditTime: 2020-11-18 08:43:32
 */
package main

import (
	"log"

	"boxin/service/api/handler"
	auth "boxin/service/auth/proto/auth"
	user "boxin/service/user/proto/user"

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

func main() {
	etcdRegister := etcd.NewRegistry(
		registry.Addrs(EtcdAddr),
	)
	// app := micro.NewService(
	// 	micro.Name(AppName),
	// 	micro.Registry(etcdRegister))
	userService := user.NewUserService("go.micro.service.user", client.DefaultClient)
	authService := auth.NewAuthService("go.micro.service.auth", client.DefaultClient)
	webHandler := gin.Default()
	service := web.NewService(
		web.Name(ServiceName),
		web.Address(":8080"),
		web.Handler(webHandler),
		web.Registry(etcdRegister),
	)
	handler.UserRouter(webHandler, userService)
	handler.AuthRouter(webHandler, authService)
	// handler.CourseRouter(webHandler,courService)
	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
