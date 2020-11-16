package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"boxin/service/user/handler"
	user "boxin/service/user/proto/user"
	repo "boxin/service/user/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const (
	ServiceName = "go.micro.service.user"
	MysqlUri = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr = "localhost:2379"
)

func main() {
	db, err := gorm.Open("mysql", MysqlUri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)

	service.Init()

	userHandler := &handler.UserHandler{UserRepository: &repo.UserRepositoryImpl{DB: db}}

	if err := user.RegisterUserServiceHandler(service.Server(), userHandler); err != nil {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}