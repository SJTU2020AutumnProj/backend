package main

import (
	"boxin/service/email/handler"
	email "boxin/service/email/proto/email"
	// 引入插件
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	// 引入公共的自定义配置函数
	"boxin/utils/tracer"
	
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"
)

/*
Configuration of Verification service
*/
const (
	ServiceName = "go.micro.service.email"
	EtcdAddr    = "localhost:2379"
	JaegerAddr  = "localhost:6831"
)

func main() {
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
		micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
	)

	service.Init()

	emaillHandler := &handler.EmailHandler{Val: 0}

	if err := email.RegisterEmailServiceHandler(service.Server(), emaillHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
