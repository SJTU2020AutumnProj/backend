/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2020-11-16 22:30:06
 * @LastEditors: Seven
 * @LastEditTime: 2020-11-25 11:04:23
 */
package main

import (
	"boxin/service/auth/handler"
	auth "boxin/service/auth/proto/auth"
	repo "boxin/service/auth/repository"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
)

const (
	ServiceName = "go.micro.service.auth"
	MysqlUri    = "root:root@(localhost:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr    = "localhost:2379"
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

	authHandler := &handler.AuthHandler{AuthRepository: &repo.AuthRepositoryImpl{DB: db}}

	if err := auth.RegisterAuthServiceHandler(service.Server(), authHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
