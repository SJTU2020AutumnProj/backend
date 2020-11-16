package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"

	user "boxin/service/user/proto/user"
	"boxin/service/api/handler"
)

const (
	ServiceName = "go.micro.api.api"
	EtcdAddr = "localhost:2379"
)

func main() {
	etcdRegister := etcd.NewRegistry(
		registry.Addrs(EtcdAddr),
	)

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