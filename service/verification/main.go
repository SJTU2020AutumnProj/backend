package main

import (
	"boxin/service/verification/handler"
	verification "boxin/service/verification/proto/verification"
	// "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	// hystrixGo "github.com/afex/hystrix-go/hystrix"
	// 自定义插件
	"boxin/service/verification/wrapper/breaker/hystrix"
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
)

func main() {
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
		micro.WrapClient(
            // 引入hystrix包装器
            hystrix.NewClientWrapper(),
        ),
	)

	// // 自定义全局默认超时时间和最大并发数
    // hystrixGo.DefaultSleepWindow = 200
	// hystrixGo.DefaultMaxConcurrent = 30
	
	// // 针对指定服务接口使用不同熔断配置
    // // 第一个参数name=服务名.接口.方法名，这并不是固定写法，而是因为官方plugin默认用这种方式拼接命令name
    // // 之后我们自定义wrapper也同样使用了这种格式
    // // 如果你采用了不同的name定义方式则以你的自定义格式为准
    // hystrixGo.ConfigureCommand("go.micro.service.email.EmailService.SendEmail",
    //     hystrixGo.CommandConfig{
    //         MaxConcurrentRequests: 50,
    //         Timeout:               20000,
    //     })

	service.Init()

	verificationHandler := &handler.VerificationHandler{VerificationRepository: &repo.VerificationRepositoryImpl{CONN: c}}

	if err := verification.RegisterVerificationServiceHandler(service.Server(), verificationHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
