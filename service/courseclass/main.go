package main

import (
	"boxin/service/courseclass/handler"
	courseclass "boxin/service/courseclass/proto/courseclass"
	repo "boxin/service/courseclass/repository"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
)

const (
	ServiceName = "go.micro.service.courseclass"
	MysqlUri    = "root:root@(host.docker.internal:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr    = "host.docker.internal:2379"
)

func main() {
	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
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

	courseHandler := &handler.CourseClassHandler{CourseClassRepository: &repo.CourseClassRepositoryImpl{DB: db}}

	if err := courseclass.RegisterCourseClassServiceHandler(service.Server(), courseHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
