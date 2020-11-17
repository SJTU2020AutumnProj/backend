package main

import (
	"log"

	"boxin/service/api/handler"
	user "boxin/service/user/proto/user"

	"github.com/gin-gonic/gin"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

const (
	ServiceName = "go.micro.api.api"
	EtcdAddr    = "localhost:2379"
)

func main() {
	etcdRegister := etcd.NewRegistry(
		registry.Addrs(EtcdAddr),
	)
	app := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcdRegister))

	userService := user.NewUserService("go.micro.service.user", app.Client())

	webHandler := gin.Default()
	service := web.NewService(
		web.Name(ServiceName),
		web.Address(":8080"),
		web.Handler(webHandler),
		web.Registry(etcdRegister),
	)
	handler.UserRouter(webHandler, userService)

	service.Init()
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
