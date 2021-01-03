package main

import (
	homework "boxin/service/homework/proto/homework"

	"bytes"
	"log"
	"context"
	"time"
	"crypto/rand"
	"math/big"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// Configuration
const (
	ServiceName = "go.micro.client.message"
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
	courseID := int32(1)
	userID := int32(1)
	startTime := time.Now().Unix()
	endTime, _ := time.Parse("2006-01-02 15:04:05", "2021-07-27 08:46:15")
	endTimeUnix := endTime.Unix()
	title := createRandomString(5)
	state := int32(1)
	description := createRandomString(20)
	content := createRandomString(200)
	assignHomeworkTest(homeworkService, courseID, userID, startTime, endTimeUnix, title, state, description, content)
}

func createRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}

func assignHomeworkTest(
	homeworkService homework.HomeworkService,
	courseID int32,
	userID int32,
	startTime int64,
	endTime int64,
	title string,
	state int32,
	description string,
	content string,
) {
	resp, err := homeworkService.AssignHomework(
		context.Background(),
		&homework.AssignHomeworkParam{
			CourseID: courseID,
			UserID: userID,
			StartTime: startTime,
			EndTime: endTime,
			Title: title,
			State: state,
			Description: description,
			Content: content,
		},
	)
	if nil != err {
		log.Println("assignHomeworkTest error ", err)
	} else {
		log.Println("assignHomeworkTest success ", resp)
	}
}
