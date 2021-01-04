package main

import (
	"boxin/service/auth/handler"
	auth "boxin/service/auth/proto/auth"
	repo "boxin/service/auth/repository"
	redis "boxin/service/auth/redis"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	redigo "github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
)

// Configuration
const (
	ServiceName = "go.micro.service.auth"
	MysqlUri    = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	RedisHost   = "127.0.0.1"
	RedisPort   = 6379
	EtcdAddr    = "localhost:2379"
)

func main() {
	db, err := gorm.Open("mysql", MysqlUri)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c, err := redigo.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	defer c.Close()

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)

	service.Init()

	authHandler := &handler.AuthHandler{
		AuthRepository: &repo.AuthRepositoryImpl{DB: db},
		AuthRedis: &redis.AuthRedisImpl{CONN: c},
	}

	if err := auth.RegisterAuthServiceHandler(service.Server(), authHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
