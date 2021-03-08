package main

import (
	mongoDB "boxin/service/news/mongoDB"
	repo "boxin/service/news/repository"
	news "boxin/service/news/proto/news"
	"boxin/service/news/handler"
	"context"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration
const (
	ServiceName = "go.micro.service.news"
	MysqlURI    = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr    = "localhost:2379"
	MongoURI    = "mongodb://localhost:27017"
	NatsURI     = "nats://localhost:4222"
)

func main() {

	//连接mysql数据库
	db, err := gorm.Open("mysql", MysqlURI)
	if nil != err {
		panic(err)
	}
	log.Println("Connected to Mysql")
	defer db.Close()

	//连接mongoDB数据库

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(MongoURI)
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
	log.Println("Connected to MongoDB")
	defer client.Disconnect(context.Background())
	collection := client.Database("jub").Collection("message")

	// New Service
	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)

	// Initialize service
	service.Init()


	// Register Handler
	newsHandler := &handler.NewsHandler{
		NewsMongo: &mongoDB.NewsMongoImpl{CL:collection},
		NewsRepository:&repo.NewsRepositoryImpl{DB: db},
	}
	if err := news.RegisterNewsServiceHandler(service.Server(), newsHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
