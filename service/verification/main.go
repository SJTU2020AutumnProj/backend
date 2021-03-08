package main

import (
	"boxin/service/verification/handler"
	verification "boxin/service/verification/proto/verification"
	"time"

	// // 自定义插件
	// "boxin/service/verification/wrapper/breaker/hystrix"
	// 引入插件
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	// 引入公共的自定义配置函数
	"boxin/utils/tracer"

	repo "boxin/service/verification/repository"
	"log"

	redigo "github.com/garyburd/redigo/redis"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
)

/*
Configuration of Verification service
*/
const (
	ServiceName = "go.micro.service.verification"
	RedisHost   = "127.0.0.1"
	RedisPort   = 6379
	EtcdAddr    = "localhost:2379"
	JaegerAddr  = "localhost:6831"
)

func main() {
	settimeout := redigo.DialConnectTimeout(60 * time.Minute)
	c, err := redigo.Dial("tcp", "127.0.0.1:6379", settimeout)
	// c, err := redigo.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		panic(err)
	}
	log.Println("Verification service connected to redis")
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

	verificationHandler := &handler.VerificationHandler{VerificationRepository: &repo.VerificationRepositoryImpl{CONN: c}}

	if err := verification.RegisterVerificationServiceHandler(service.Server(), verificationHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
