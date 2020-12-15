package main

import (
	"boxin/service/answer/handler"
	answer "boxin/service/answer/proto/answer"
	repo "boxin/service/answer/repository"
	mongoDB "boxin/service/answer/mongoDB"

	"context"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ServiceName = "go.micro.service.answer"
	MysqlUri    = "root:root@(127.0.0.1:3306)/jub?charset=utf8mb4&parseTime=True&loc=Local"
	EtcdAddr    = "localhost:2379"
)

func main() {
	//连接mysql数据库

	db, err := gorm.Open("mysql", MysqlUri)
	if nil != err {
		panic(err)
	}
	defer db.Close()

	//连接mongoDB数据库

	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
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
	fmt.Println("Connected to MongoDB!")

	collection := client.Database("jub").Collection("answer")


	//启动服务
	service := micro.NewService(
		micro.Name(ServiceName),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)

	service.Init()

	answerHandler := &handler.AnswerHandler{AnswerMongo: &mongoDB.AnswerMongoImpl{CL:collection},AnswerRepository:&repo.AnswerRepositoryImpl{DB: db}}

	if err := answer.RegisterAnswerServiceHandler(service.Server(), answerHandler); nil != err {
		log.Fatal(errors.WithMessage(err, "register server"))
	}

	if err := service.Run(); nil != err {
		log.Fatal(errors.WithMessage(err, "run server"))
	}
}
