/*
 * @Description:
 * @Version: 1.0
 * @Author: Zhang AO
 * @studentID: 518021910368
 * @School: SJTU
 * @Date: 2021-01-08 00:06:29
 * @LastEditors: Seven
 * @LastEditTime: 2021-01-08 09:57:23
 */
package main

import (
	"boxin/service/auth/handler"
	auth "boxin/service/auth/proto/auth"
	redis "boxin/service/auth/redis"
	repo "boxin/service/auth/repository"
	"log"
	"time"

	// 引入插件
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	// 引入公共的自定义配置函数
	"boxin/utils/tracer"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	JaegerAddr  = "localhost:6831"
)

func main() {
	db, err := gorm.Open("mysql", MysqlUri)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	settimeout := redigo.DialConnectTimeout(60 * time.Minute)

	c, err := redigo.Dial("tcp", "127.0.0.1:6379", settimeout)
	if err != nil {
		panic(err)
	}
	log.Println("Auth service connected to redis")
	defer c.Close()

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(ServiceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
		micro.WrapClient(
			// // 引入hystrix包装器
			// hystrix.NewClientWrapper(),
			// 配置链路追踪为jaeger
			opentracing.NewClientWrapper(jaegerTracer),
		),
		// 配置链路追踪为jaeger
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
	)

	service.Init()

	authHandler := &handler.AuthHandler{
		AuthRepository: &repo.AuthRepositoryImpl{DB: db},
		AuthRedis:      &redis.AuthRedisImpl{CONN: c},
	}

	if err := auth.RegisterAuthServiceHandler(service.Server(), authHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
