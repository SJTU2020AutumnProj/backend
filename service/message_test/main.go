package main

import (
	homework "boxin/service/homework/proto/homework"
	message "boxin/service/message/proto/message"
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
	messageService := message.NewMessageService("go.micro.service.message", server.Client())
	courseID := int32(1)
	teacherID := int32(1)
	studentID := int32(1)
	startTime := time.Now().Unix()
	endTime, _ := time.Parse("2006-01-02 15:04:05", "2021-07-27 08:46:15")
	endTimeUnix := endTime.Unix()
	title := createRandomString(5)
	state := int32(1)
	description := createRandomString(20)
	content := createRandomString(200)
	note := createRandomString(50)
	homeworkID := assignHomeworkTest(homeworkService, courseID, teacherID, startTime, endTimeUnix, title, state, description, content, note)
	answerID := postHomeworkAnswerTest(homeworkService, homeworkID, teacherID, time.Now().Unix(), content, note)
	homeworkAnswerPubTest(homeworkService, homeworkID, answerID, teacherID, courseID, title, time.Now().Unix())
	getMessageByCourseIDTest(messageService, courseID)
	getMessageByUserIDTest(messageService, studentID)
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
	note string,
) int32 {
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
			Note: note,
		},
	)
	if nil != err {
		log.Println("assignHomeworkTest error ", err)
		return -1
	}
	log.Println("assignHomeworkTest success ", resp)
	return resp.HomeworkID
}

func postHomeworkAnswerTest(
	homeworkService homework.HomeworkService,
	homeworkID int32,
	userID int32,
	commitTime int64,
	content string,
	note string,
) int32 {
	resp, err := homeworkService.PostHomeworkAnswer(
		context.Background(),
		&homework.PostParam{
			HomeworkID: homeworkID,
			UserID: userID,
			CommitTime: commitTime,
			Content: content,
			Note: note,
		},
	)
	if nil != err {
		log.Println("postHomeworkAnswerTest error ", err)
		return -1
	}
	log.Println("postHomeworkAnswerTest success ", resp)
	return resp.AnswerID
}

func homeworkAnswerPubTest(
	homeworkService homework.HomeworkService,
	homeworkID int32,
	answerID int32,
	teacherID int32,
	courseID int32,
	title string,
	pubTime int64,
) {
	resp, err := homeworkService.ReleaseHomeworkAnswer(
		context.Background(),
		&homework.ReleaseParam{
			HomeworkID: homeworkID,
			AnswerID: answerID,
			TeacherID: teacherID,
			CourseID: courseID,
			PubTime: pubTime,
		},
	)
	if nil != err {
		log.Println("homeworkAnswerPubTest error ", err)
	} else {
		log.Println("homeworkAnswerPubTest success ", resp)
	}
}

func releaseChcekTest(
	homeworkService homework.HomeworkService,
	homeworkID int32,
	teacherID int32,
	courseID int32,
	releaseTime int64,
) {
	resp, err := homeworkService.ReleaseCheck(
		context.Background(),
		&homework.ReleaseCheckParam{
			HomeworkID: homeworkID,
			TeacherID: teacherID,
			CourseID: courseID,
			ReleaseTime: releaseTime,
		},
	)
	if nil != err {
		log.Println("homeworkAnswerPubTest error ", err)
	} else {
		log.Println("homeworkAnswerPubTest success ", resp)
	}
}

func getMessageByUserIDTest(
	messageService message.MessageService,
	userID int32,
) {
	resp, err := messageService.GetMessageByUserID (
		context.Background(),
		&message.GetMessageByUserIDParam {
			UserID: userID,
		},
	)
	if nil != err {
		log.Println("getMessageByUserIDTest error ", err)
	} else {
		log.Println("getMessageByUserIDTest success ", resp)
	}
}

func getMessageByCourseIDTest(
	messageService message.MessageService,
	courseID int32,
) {
	resp, err := messageService.GetMessageByCourseID (
		context.Background(),
		&message.GetMessageByCourseIDParam {
			CourseID: courseID,
		},
	)
	if nil != err {
		log.Println("getMessageByCourseIDTest error ", err)
	} else {
		log.Println("getMessageByCourseIDTest success ", resp)
	}
}
