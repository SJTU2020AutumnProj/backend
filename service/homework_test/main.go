package main

import (
	homework "boxin/service/homework/proto/homework"
	// news "boxin/service/news/proto/news"
	// "bytes"
	// "log"
	// "time"
	// "crypto/rand"
	// "math/big"
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// Configuration
const (
	ServiceName = "go.micro.client.homework"
	EtcdAddr    = "localhost:2379"
)

func main() {
	server := micro.NewService(
		micro.Name(ServiceName),
		micro.Registry(etcd.NewRegistry(
			registry.Addrs(EtcdAddr),
		)),
	)
	server.Init()
	homeworkService := homework.NewHomeworkService("go.micro.service.homework", server.Client())
	// var homeworkID *homework.HomeworkID
	// homeworkID.HomeworkID = 1
	// fmt.Println(1)


	resp,err:=homeworkService.GetUserByHomeworkID(context.Background(),&homework.HomeworkID{HomeworkID:1})

	// resp,err:=homeworkService.SearchHomework(context.Background(),&homework.HomeworkID{HomeworkID:1})
	fmt.Println(resp,err)
}