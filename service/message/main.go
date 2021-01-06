package main

import (
	mongoDB "boxin/service/message/mongoDB"
	repo "boxin/service/message/repository"
	message "boxin/service/message/proto/message"
	"boxin/service/message/subscriber"
	"boxin/service/message/handler"
	"context"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/broker/nats"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Configuration
const (
	ServiceName = "go.micro.service.message"
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
		micro.Broker(nats.NewBroker(
			broker.Addrs(NatsURI),
		)),
	)

	// Initialize service
	service.Init()

	// Register Subscriber
	subscriberHandler := &subscriber.MessageSub{
		MessageRepository: &repo.MessageRepositoryImpl{
			DB: db,
		},
		MongoDB: &mongoDB.MessageMongoImpl{
			CL: collection,
		},
	}
	// 这里的topic注意与homework注册的要一致
	if err := micro.RegisterSubscriber("go.micro.service.homework.assigned", service.Server(), subscriberHandler.Assigned); err != nil {
		log.Fatal(errors.WithMessage(err, "subscribe"))
	}
	if err := micro.RegisterSubscriber("go.micro.service.homework.published", service.Server(), subscriberHandler.PostAnswer); err != nil {
		log.Fatal(errors.WithMessage(err, "subscribe"))
	}
	if err := micro.RegisterSubscriber("go.micro.service.homework.checkReleased", service.Server(), subscriberHandler.ReleaseCheck); err != nil {
		log.Fatal(errors.WithMessage(err, "subscribe"))
	}

	// Register Handler
	messageHandler := &handler.MessageHandler{
		MessageMongo: &mongoDB.MessageMongoImpl{CL:collection},
		MessageRepository:&repo.MessageRepositoryImpl{DB: db},
	}
	if err := message.RegisterMessageServiceHandler(service.Server(), messageHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}


	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
