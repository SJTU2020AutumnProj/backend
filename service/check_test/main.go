package main

import (
	check "boxin/service/check/proto/check"

	"bytes"
	"context"
	"crypto/rand"
	"log"
	"math/big"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

// Configuration
const (
	ServiceName = "go.micro.client.check"
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
	checkService := check.NewCheckService("go.micro.service.check", server.Client())
	answerID := int32(0)
	homeworkID := int32(1)
	teacherID := int32(0)
	studentID := int32(1)
	checkTime := time.Now().Unix()
	description := createRandomString(100)
	comment := createRandomString(50)
	score := int32(100)
	addCheckTest(checkService, answerID, homeworkID, teacherID, studentID, checkTime, description, comment, score)
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

func addCheckTest(
	checkService check.CheckService,
	answerID int32,
	homeworkID int32,
	teacherID int32,
	studentID int32,
	checkTime int64,
	description string,
	comment string,
	score int32,
) {
	resp, err := checkService.CreateCheck(
		context.Background(),
		&check.CreateCheckParam{
			AnswerID:    answerID,
			HomeworkID:  homeworkID,
			TeacherID:   teacherID,
			StudentID:   studentID,
			CheckTime:   checkTime,
			Description: description,
			Comment:     comment,
			Score:       score,
		},
	)
	if resp.Status != 0 {
		log.Println("addChcekTest error ", err)
	} else {
		log.Println("addChcekTest success ", resp)
	}

}
