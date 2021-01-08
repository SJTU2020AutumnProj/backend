package main

import (
	"boxin/service/check/handler"
	mongoDB "boxin/service/check/mongoDB"
	check "boxin/service/check/proto/check"
	repo "boxin/service/check/repository"

	"context"
	"fmt"
	"log"

	// 引入插件
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	// 引入公共的自定义配置函数
	"boxin/utils/tracer"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration
const (
	ServiceName = "go.micro.service.check"
	MysqlURI    = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr    = "localhost:2379"
	MongoDBURI  = "mongodb://localhost:27017"
	JaegerAddr  = "localhost:6831"
)

func main() {

	//连接mysql数据库
	db, err := gorm.Open("mysql", MysqlURI)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	//连接mongoDB数据库
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(MongoDBURI)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	collection := client.Database("jub").Collection("check")
	defer client.Disconnect(context.Background())

	// 配置jaeger连接
	jaegerTracer, closer, err := tracer.NewJaegerTracer(ServiceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()

	//启动服务
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

	checkHandler := &handler.CheckHandler{CheckMongo: &mongoDB.CheckMongoImpl{CL: collection}, CheckRepository: &repo.CheckRepositoryImpl{DB: db}}
	if err := check.RegisterCheckServiceHandler(service.Server(), checkHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}
	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
