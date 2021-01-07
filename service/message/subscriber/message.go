package subscriber

import (
	homework "boxin/service/homework/proto/homework"
	mongo "boxin/service/message/mongoDB"
	repo "boxin/service/message/repository"
	"context"
	"log"
	"time"
)

// MessageSub struct
type MessageSub struct {
	MessageRepository repo.MessageRepository
	MongoDB           mongo.MessageMongo
}

// Assigned add an assigned homework message
func (sub *MessageSub) Assigned(ctx context.Context, homework *homework.AssignedHomework) error {
	log.Println("Subscriber Assigned received message: ", homework)
	now := time.Unix(time.Now().Unix(), 0)
	message, err := sub.MessageRepository.AddMessage(
		ctx,
		repo.Message{
			MessageTime: now,
			MessageType: 0,
			CourseID:    homework.CourseID,
			Title:       "作业 " + homework.Title + " 已经发布",
			State:       0,
		},
	)
	if nil != err {
		log.Println("Subscriber Assigned add new message in Mysql error ", err)
		return err
	}
	content := "作业 " + homework.Title + " 已经发布，" + "截止时间为 " + time.Unix(homework.EndTime, 0).Format("2006-01-02 15:04:05")
	err = sub.MongoDB.AddMessage(
		ctx,
		mongo.Message{
			MessageID: message.MessageID,
			Content:   content,
		},
	)
	if nil != err {
		log.Println("Subscriber Assigned add new message in MongoDB error ", err)
		return err
	}
	return nil
}


func (sub *MessageSub) Post(ctx context.Context, message *homework.PostHomeworkAnswer) error {
	log.Println("Subscriber Post received message: ", message)
	now := time.Unix(time.Now().Unix(), 0)
	res, err := sub.MessageRepository.AddMessage(
		ctx,
		repo.Message{
			MessageTime: now,
			MessageType: 0,
			CourseID:    message.CourseID,
			Title:       "作业 " + message.Title + " 答案已经发布",
			State:       0,
		},
	)
	if nil != err {
		log.Println("Subscriber Post add new message in Mysql error ", err)
		return err
	}
	content := "作业 " + message.Title + " 答案已经发布，请前往相应页面查看"
	err = sub.MongoDB.AddMessage(
		ctx,
		mongo.Message{
			MessageID: res.MessageID,
			Content:   content,
		},
	)
	if nil != err {
		log.Println("Subscriber Post add new message in MongoDB error ", err)
		return err
	}
	return nil
}